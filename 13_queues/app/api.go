package app

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type apiConfig struct {
	Port int
}

type api struct {
	config *config
}

func NewAPI(config *config) *api {
	return &api{
		config: config,
	}
}

func (a *api) Run() error {
	r := mux.NewRouter()

	ctx := context.Background()

	redis, err := NewRedisWrapper(ctx, a.config.Redis)
	if err != nil {
		return fmt.Errorf("creating new redis wrapper: %w", err)
	}

	redisHandlers := NewRedisHandlers(redis)

	r.HandleFunc("/redis/{queue}", redisHandlers.Pop).Methods(http.MethodGet)
	r.HandleFunc("/redis/{queue}", redisHandlers.Push).Methods(http.MethodPost)

	beanstalk, err := NewBeanstalkWrapper(a.config.Beanstalk)
	if err != nil {
		return fmt.Errorf("creating new beanstalk wrapper: %w", err)
	}

	beanstalkHandlers := NewBeanstalkHandlers(beanstalk)

	r.HandleFunc("/beanstalk/{queue}", beanstalkHandlers.Pop).Methods(http.MethodGet)
	r.HandleFunc("/beanstalk/{queue}", beanstalkHandlers.Push).Methods(http.MethodPost)

	log.Printf("starting api on port %d\n", a.config.API.Port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", a.config.API.Port), r))

	return nil
}
