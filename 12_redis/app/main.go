package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/mediocregopher/radix/v4"
)

func main() {
	router := mux.NewRouter()

	db := &inMemDB{
		data: map[string]string{
			"key1": "value1",
			"key2": "value2",
			"key3": "value3",
		},
	}

	cache, err := newCache(context.Background(), *newRedisConfig())
	if err != nil {
		log.Fatalf("failed to create cache: %v", err)
	}

	api := newAPI(db, cache)

	router.HandleFunc("/simple/data/{key}", api.getData).Methods(http.MethodGet)
	router.HandleFunc("/data/{key}", api.getDataWithProbabilisticEarlyExpiration).Methods(http.MethodGet)
	router.HandleFunc("/ping", api.ping).Methods(http.MethodGet)

	log.Fatalf("serving failed: %w", http.ListenAndServe(":80", router))
}

type api struct {
	db    *inMemDB
	cache *cache
}

func newAPI(db *inMemDB, cache *cache) *api {
	return &api{
		db:    db,
		cache: cache,
	}
}

func (a *api) getData(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	k := vars["key"]

	v, ok := a.cache.get(r.Context(), k)
	if !ok {
		log.Printf("cache miss: %s", k)
		v, ok = a.db.get(k)
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		err := a.cache.set(r.Context(), k, v)
		if err != nil {
			log.Printf("failed to set cache: %v", err)
		}
	} else {
		log.Printf("cache hit: %s", k)
	}

	w.Write([]byte(v))
}

func (a *api) getDataWithProbabilisticEarlyExpiration(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	k := vars["key"]

	d, ok := a.cache.getWithDeltaAndTTL(r.Context(), k)
	if !ok || d.delta > d.ttl {
		log.Printf("updating %s in cache", k)
		v, ok := a.db.get(k)
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		d = &data{
			value: v,
			delta: 10,
			ttl:   100,
		}

		err := a.cache.setWithDeltaAndTTL(r.Context(), k, d)
		if err != nil {
			log.Printf("failed to set cache: %v", err)
		}
	}

	w.Write([]byte(d.value))
}

func (a *api) ping(w http.ResponseWriter, r *http.Request) {
	if !a.cache.ping(r.Context()) {
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// cache

type redisConfig struct {
	SentinelName  string
	SentinelAddrs []string
	RedisAddr     string
}

func newRedisConfig() *redisConfig {
	v := os.Getenv("SentinelAddrs")

	return &redisConfig{
		SentinelAddrs: strings.Split(v, ","),
		SentinelName:  os.Getenv("SentinelName"),
		RedisAddr:     os.Getenv("RedisAddr"),
	}
}

type cache struct {
	c *radix.Sentinel
	// c radix.Client
}

func newCache(ctx context.Context, config redisConfig) (*cache, error) {
	c, err := (radix.SentinelConfig{}).New(ctx, config.SentinelName, config.SentinelAddrs)
	// c, err := (radix.PoolConfig{}).New(ctx, "tcp", config.RedisAddr)
	if err != nil {
		return nil, err
	}

	return &cache{
		c: c,
	}, nil
}

func (c *cache) get(ctx context.Context, key string) (string, bool) {
	var v string
	err := c.c.Do(ctx, radix.Cmd(&v, "GET", key))
	if err != nil || v == "" {
		return "", false
	}

	return v, true
}

type data struct {
	value string
	delta int
	ttl   int
}

func (c *cache) getWithDeltaAndTTL(ctx context.Context, key string) (*data, bool) {
	var v []string
	err := c.c.Do(ctx, radix.Cmd(&v, "HMGET", key, "value", "delta"))
	if err != nil || len(v) == 0 || v[0] == "" {
		log.Printf("key %s not found\n", key)
		return nil, false
	}

	dlt, err := strconv.Atoi(v[1])
	if err != nil {
		log.Println("delta is not a number")
		return nil, false
	}

	d := data{
		value: v[0],
		delta: dlt,
	}

	var ttl int
	err = c.c.Do(ctx, radix.Cmd(&ttl, "TTL", key))
	if err != nil || ttl <= 0 {
		log.Printf("key %s ttl returned %d\n", key, ttl)
		return nil, false
	}

	d.ttl = ttl

	log.Printf("cache hit: key=%s ttl=%d", key, ttl)

	return &d, true
}

func (c *cache) set(ctx context.Context, key, value string) error {
	return c.c.Do(ctx, radix.Cmd(nil, "SET", key, value))
}

func (c *cache) setWithDeltaAndTTL(ctx context.Context, key string, d *data) error {
	if d == nil {
		return fmt.Errorf("cannot set empty data to %s", key)
	}

	err := c.c.Do(ctx, radix.Cmd(nil, "HSET", key, "value", d.value, "delta", fmt.Sprintf("%d", d.delta)))
	if err != nil {
		return fmt.Errorf("HSET %s: %w", key, err)
	}

	err = c.c.Do(ctx, radix.Cmd(nil, "EXPIRE", key, fmt.Sprintf("%d", d.ttl)))
	if err != nil {
		return fmt.Errorf("EXPIRE %s: %w", key, err)
	}

	return nil
}

func (c *cache) ping(ctx context.Context) bool {
	var v string
	err := c.c.Do(ctx, radix.Cmd(&v, "PING"))
	if err != nil {
		return false
	}

	return v == "PONG"
}

// db

type inMemDB struct {
	data map[string]string
}

func (db *inMemDB) get(key string) (string, bool) {
	v, exist := db.data[key]
	return v, exist
}
