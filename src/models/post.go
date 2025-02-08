package models

import "time"

type Post struct {
	ID        uint `gorm:"primaryKey"`
	authorId  uint `gorm:"foreignKey:Users"`
	title     string
	body      string
	createdAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
