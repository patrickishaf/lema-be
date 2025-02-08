package models

type User struct {
	ID       uint `gorm:"primaryKey"`
	name     string
	username string
	email    string
	address  string
}
