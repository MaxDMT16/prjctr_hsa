package main

import (
	"fmt"
	app "prjctr/md/13_queues"
)

func main() {
	config, err := app.NewConfig()
	if err != nil {
		panic(fmt.Errorf("creating new config: %w", err))
	}

	api := app.NewAPI(config)

	api.Run()
}