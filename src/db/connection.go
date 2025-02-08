package db

import (
	"log"

	"github.com/patrickishaf/lema-be/src/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitializeDb() {
	database, err := gorm.Open(sqlite.Open("main.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to db")
	}

	db = database
	dbError := db.AutoMigrate(&models.Post{})

	if dbError != nil {
		log.Println("failed to migrate database", dbError)
	}
}
