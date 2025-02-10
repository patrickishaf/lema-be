package models

import "time"

type Post struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	AuthorId  uint      `json:"author_id" gorm:"foreignKey:Users"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
