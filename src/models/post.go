package models

import "time"

type Post struct {
	ID        uint `gorm:"primaryKey"`
	AuthorId  uint `gorm:"foreignKey:Users"`
	Title     string
	Body      string
	CreatedAt time.Time `gorm:"autoCreateTime"`
	CpdatedAt time.Time `gorm:"autoUpdateTime"`
}
