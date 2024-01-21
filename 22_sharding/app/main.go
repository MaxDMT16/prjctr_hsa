package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	db := DB()

	var books []Book
	tx := db.Find(&books)
	if tx.Error != nil {
		panic(tx.Error)
	}

	fmt.Println(books)
}

type Book struct {
	ID int `gorm:"primaryKey"`
	CategoryID int `gorm:"column:category_id, not null"`
	Autor string `gorm:"not null"`
	Title string `gorm:"not null"`
	Year int `gorm:"not null"`
}


// host=localhost user=prjctr password=test dbname=prjctr port=5432 sslmode=disable
func DB() *gorm.DB {
	// dsn := os.Getenv("DSN")
	dsn := "host=localhost user=prjctr password=test dbname=prjctr port=5432 sslmode=disable"
	if dsn == "" {
		panic("DSN env var is required")
	}

	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic(err)
	}

	return db
}
