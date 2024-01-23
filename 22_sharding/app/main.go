package main

import (
	"fmt"
	"os"
	"runtime"
	"sync"
	"time"

	"github.com/brianvoe/gofakeit"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const BooksCount = 1_000_000
// const BooksCount = 1_00

func main() {
	db := DB()

	writeStartTime := time.Now()

	writeBooks(db, BooksCount)

	fmt.Printf("write %d books took %fs\n", BooksCount, time.Since(writeStartTime).Seconds())
}

func writeBooks(db *gorm.DB, booksCount int) {
	var wg sync.WaitGroup

	goroutinesCount := runtime.NumCPU()
	wg.Add(goroutinesCount)
	
	batchSize := booksCount / goroutinesCount

	for i := 0; i < goroutinesCount; i++ {
		go func() {
			for j := 0; j <= batchSize; j++ {
				b := generateFakeBook()

				err := CreateBooks(db, b)
				if err != nil {
					panic(err)
				}
			}

			wg.Done()
		}()
	}

	fmt.Println("Work has been launched...")

	wg.Wait()

	fmt.Printf("%d books have been created\n", booksCount)
}

type Book struct {
	ID         int    `gorm:"primaryKey"`
	CategoryID int    `gorm:"column:category_id;not null"`
	Author      string `gorm:"not null"`
	Title      string `gorm:"not null"`
	Year       int    `gorm:"not null"`
}

func generateFakeBook() Book {
	return Book{
		ID:         gofakeit.Number(1, 1_000_000),
		CategoryID: gofakeit.Number(1, 10),
		Author:      gofakeit.Name(),
		Title:      gofakeit.Sentence(3),
		Year:       gofakeit.Year(),
	}
}

func DB() *gorm.DB {
	dsn := os.Getenv("APP_DSN")

	if dsn == "" {
		panic("APP_DSN env var is required")
	}

	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic(err)
	}

	return db
}

func CreateBooks(db *gorm.DB, b Book) error {
	statement := `INSERT INTO books (id, category_id, author, title, year) VALUES (?, ?, ?, ?, ?)`

	tx := db.Exec(statement, b.ID, b.CategoryID, b.Author, b.Title, b.Year)
	if tx.Error != nil {
		return fmt.Errorf("creating book in DB: %w", tx.Error)
	}

	return nil
}
