package app

import (
	"time"

	"github.com/brianvoe/gofakeit/v6"
)

type User struct {
	ID    int    `json:"-" gorm:"primaryKey, autoIncrement"`
	Name  string `json:"name" gorm:"name"`
	Email string `json:"email" gorm:"email"`
	Age int    `json:"age" gorm:"age"`
	DateOfBirth string `json:"date_of_birth" gorm:"date_of_birth"`
}

func NewRandomUser() *User {
	minDateOfBirth := time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)
	maxDateOfBirth := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)

	return &User{
		Name:  gofakeit.Name(),
		Email: gofakeit.Email(),
		Age:   gofakeit.Number(18, 80),
		DateOfBirth: gofakeit.DateRange(minDateOfBirth, maxDateOfBirth).Format(time.DateOnly),
	}
}