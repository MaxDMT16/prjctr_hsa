package app

type User struct {
	ID    int    `json:"-" gorm:"primaryKey, autoIncrement"`
	Name  string `json:"name" gorm:"name"`
	Email string `json:"email" gorm:"email"`
	Age int    `json:"age" gorm:"age"`
	DateOfBirth string `json:"date_of_birth" gorm:"date_of_birth"`
}
