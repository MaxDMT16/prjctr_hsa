package main

import (
	"fmt"
	"log"
	app "prjctr/md/13_queues"
)

func main() {
	config, err := app.NewConfig()
	if err != nil {
		panic(fmt.Errorf("creating new config: %w", err))
	}

	api := app.NewAPI(config)

	log.Fatal(api.Run())
}