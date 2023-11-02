package app

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewStorage() *gorm.DB {
	dsn := os.Getenv("POSTGRES_DSN")
	// dsn := "host=localhost user=postgres password=mysecretpassword dbname=main port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("creating db connection failed \n connection string: %v \n err: %v", dsn, err))
	}

	logrus.Debug("creating new db connection")
	return db
}