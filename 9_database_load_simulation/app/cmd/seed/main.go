package main

import (
	"sync"
	"time"

	"github.com/brianvoe/gofakeit/v6"
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
	goroutines := 500
	desiredUsersCount := 40_000_000
	usersPerGoroutine := desiredUsersCount / goroutines
	usersInTx := 3_000

	start := time.Now()

	for i := 0; i < goroutines; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for j := 0; j < usersPerGoroutine/usersInTx; j++ {
				storage.Transaction(func(tx *gorm.DB) error {
					for k := 0; k < usersInTx; k++ {
						user := generateUser()
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

	println("elapsed time: ", elapsed)
}

func generateUser() *app.User {
	minDateOfBirth := time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)
	maxDateOfBirth := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)

	return &app.User{
		Name:        gofakeit.Name(),
		Email:       gofakeit.Email(),
		Age:         gofakeit.Number(18, 80),
		DateOfBirth: gofakeit.DateRange(minDateOfBirth, maxDateOfBirth).Format(time.DateOnly),
	}
}
