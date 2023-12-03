package app

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mediocregopher/radix/v4"
)


type redisHandlers struct {
	redis *redisWrapper
}

func NewRedisHandlers(redis *redisWrapper) *redisHandlers {
	return &redisHandlers{
		redis: redis,
	}
}

func (h *redisHandlers) Push(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	queue := vars["queue"]

	var body PushRequest
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		log.Println("invalid json body")
		
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if body.Message == "" {
		log.Println("message is required")

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	h.redis.Push(r.Context(), queue, body.Message)

	w.Write([]byte(queue))
}

func (h *redisHandlers) Pop(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	queue := vars["queue"]

	msg, err := h.redis.Pop(r.Context(), queue)
	if err != nil {
		log.Printf("pop from queue %s: %s\n", queue, err.Error())

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(PopResponse{
		Message: msg,
	})
}

type redisConfig struct {
	Address string
}

type redisWrapper struct {
	client radix.Client
}

func NewRedisWrapper(ctx context.Context, config redisConfig) (*redisWrapper, error) {
	client, err := (radix.PoolConfig{}).New(ctx, "tcp", config.Address)
	if err != nil {
		return nil, fmt.Errorf("create redis client: %w", err)
	}

	return &redisWrapper{
		client: client,
	}, nil
}

func (r *redisWrapper) Push(ctx context.Context, queue string, msg string) error {
	return r.client.Do(ctx, radix.Cmd(nil, "LPUSH", queue, msg))
}

func (r *redisWrapper) Pop(ctx context.Context, queue string) (string, error) {
	var msg string
	err := r.client.Do(ctx, radix.Cmd(&msg, "RPOP", queue))
	if err != nil {
		return "", err
	}
	return msg, nil
}
