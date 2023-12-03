package app

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mediocregopher/radix/v4"
)

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

	r.HandleFunc("/redis/{queue}/push", func(w http.ResponseWriter, r *http.Request) {
	})

	http.ListenAndServe(fmt.Sprintf(":%d", a.config.API.Port), nil)

	return nil
}

type redisHandlers struct {
	redis *radix.Client
}

func (h *redisHandlers) Push(w http.ResponseWriter, r *http.Request) {

}
