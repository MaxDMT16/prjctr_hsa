package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	beanstalk "github.com/beanstalkd/go-beanstalk"
	"github.com/gorilla/mux"
)

type beanstalkHandlers struct {
	beanstalkd *beanstalkdWrapper
}

func NewBeanstalkHandlers(beanstalkd *beanstalkdWrapper) *beanstalkHandlers {
	return &beanstalkHandlers{
		beanstalkd: beanstalkd,
	}
}

func (h *beanstalkHandlers) Push(w http.ResponseWriter, r *http.Request) {
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

	h.beanstalkd.Push(queue, body.Message)

	w.Write([]byte(queue))
}

func (h *beanstalkHandlers) Pop(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	queue := vars["queue"]

	message, err := h.beanstalkd.Pop(queue)
	if err != nil {
		log.Println("error popping message from beanstalkd")

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(PopResponse{
		Message: message,
	})
}

type beanstalkConfig struct {
	Address string
}

type beanstalkdWrapper struct {
	client *beanstalk.Conn
}

func NewBeanstalkWrapper(config beanstalkConfig) (*beanstalkdWrapper, error) {
	c, err := beanstalk.Dial("tcp", config.Address)
	if err != nil {
		return nil, fmt.Errorf("dial beanstalk: %w", err)
	}

	return &beanstalkdWrapper{
		client: c,
	}, nil
}

func (b *beanstalkdWrapper) Push(queue, message string) error {
	b.client.Tube.Name = queue

	_, err := b.client.Put([]byte(message), 1, 0, 0)
	if err != nil {
		return fmt.Errorf("pushing message to beanstalk: %w", err)
	}

	return nil
}

func (b *beanstalkdWrapper) Pop(queue string) (string, error) {
	b.client.Tube.Name = queue

	jobID, body, err := b.client.Reserve(1)
	if err != nil {
		return "", fmt.Errorf("reserving job %d from beanstalk: %w", jobID, err)
	}

	return string(body), nil
}
