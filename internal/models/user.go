package models

type User struct {
	ID           uint   `gorm:"primaryKey" json:"id"`
	Name         string `json:"name"`
	Email        string `gorm:"unique" json:"email"`
	Age          int    `json:"age"`
	PasswordHash string `json:"-"`
}
