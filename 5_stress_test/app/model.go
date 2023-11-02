package app

type SendEventModel struct {
	ID    int    `json:"-", gorm:"primaryKey, autoIncrement"`
	Name  string `json:"name", gorm:"name"`
	Value int    `json:"value", gorm:"value"`
}
