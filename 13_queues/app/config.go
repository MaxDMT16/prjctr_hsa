package app


type config struct {
	API apiConfig
	Redis redisConfig
	Beanstalk beanstalkConfig
}

func NewConfig() (*config, error) {
	// port, err := strconv.Atoi(os.Getenv("API_PORT"))
	// if err != nil {
	// 	return nil, fmt.Errorf("port value is not valid integer: %w", err)
	// }

	return &config{
		API: apiConfig{
			Port: 8989,//port,
		},
		Redis: redisConfig{
			Address: "localhost:6379", // os.Getenv("REDIS_ADDRESS"),
		},
		Beanstalk: beanstalkConfig{
			Address: "localhost:11300", // os.Getenv("BEANSTALK_ADDRESS"),
		},
	}, nil
}