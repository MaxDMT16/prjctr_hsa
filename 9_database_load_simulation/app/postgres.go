package app

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewStorage() *gorm.DB {
	// dsn := os.Getenv("POSTGRES_DSN")
	dsn := "host=127.0.0.1 user=test_user password=test_user_pass dbname=main port=5432"
	
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.New(logrus.New(), logger.Config{
			LogLevel: logger.Info,
		}),
	})
	if err != nil {
		logrus.Fatal(fmt.Sprintf("creating db connection failed \n connection string: %v \n err: %v", dsn, err))
	}

	sqlDB, err := db.DB()
	if err != nil {
		logrus.Fatal(fmt.Sprintf("getting sql db from gorm failed: %v", err))
	}

	sqlDB.SetMaxOpenConns(0) // explicilty set to unlimited

	logrus.Debug("creating new db connection")
	return db
}