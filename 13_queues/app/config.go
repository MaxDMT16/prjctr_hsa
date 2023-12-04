package app

import (
	"fmt"
	"os"
	"strconv"
)

type config struct {
	API       apiConfig
	Redis     redisConfig
	Beanstalk beanstalkConfig
}

func NewConfig() (*config, error) {
	port, err := strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		return nil, fmt.Errorf("port value is not valid integer: %w", err)
	}

	jobTTR, err := strconv.Atoi(os.Getenv("BEANSTALK_JOB_TTR"))
	if err != nil {
		return nil, fmt.Errorf("job TTR value is not valid integer: %w", err)
	}

	jobTimeout, err := strconv.Atoi(os.Getenv("BEANSTALK_JOB_TIMEOUT"))
	if err != nil {
		return nil, fmt.Errorf("job timeout value is not valid integer: %w", err)
	}

	return &config{
		API: apiConfig{
			Port: port, // 8989,
		},
		Redis: redisConfig{
			Address: os.Getenv("REDIS_ADDRESS"), // "localhost:6379",
		},
		Beanstalk: beanstalkConfig{
			Address:    os.Getenv("BEANSTALK_ADDRESS"), // "localhost:11300",
			JobTTR:     jobTTR,                         // 60,
			JobTimeout: jobTimeout,                     //10
		},
	}, nil
}
