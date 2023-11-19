package main

import (
	"sync"
	"time"

	"gorm.io/gorm"

	app "prjctr/md/9_database_load_simulation"
)

func main() {
	storage := app.NewStorage()
	err := storage.AutoMigrate(&app.User{})
	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup
	goroutines := 10
	desiredUsersCount := 40_000
	usersPerGoroutine := desiredUsersCount / goroutines
	usersInTx := 250

	start := time.Now()

	for i := 0; i < goroutines; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for j := 0; j < usersPerGoroutine/usersInTx; j++ {
				storage.Transaction(func(tx *gorm.DB) error {
					for k := 0; k < usersInTx; k++ {
						user := app.NewRandomUser()
						err := tx.Create(user).Error
						if err != nil {
							return err
						}
					}

					return nil
				})
			}
		}()
	}

	wg.Wait()

	elapsed := time.Since(start)

	println("elapsed time in sec: ", elapsed.Seconds())
}
