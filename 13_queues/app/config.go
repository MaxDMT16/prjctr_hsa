package app

import (
	"fmt"
	"os"
	"strconv"
)


type config struct {
	API API
}

type API struct {
	Port int
}

func NewConfig() (*config, error) {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		return nil, fmt.Errorf("port value is not valid integer: %w", err)
	}

	return &config{
		API: API{
			Port: port,
		},
	}, nil
}