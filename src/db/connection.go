package db

import (
	"log"

	"github.com/patrickishaf/lema-be/src/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitializeDb() {
	database, err := gorm.Open(sqlite.Open("../main.sqlite3"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to db")
	}

	db = database
	dbError := db.AutoMigrate(&models.Post{}, &models.User{})

	if dbError != nil {
		log.Println("failed to migrate database", dbError)
	}

	// insertError := insertDummyUsers()
	// if insertError != nil {
	// 	log.Println("failed to indsert dummy users", insertError)
	// }
	// insertError := insertDummyPosts()
	// if insertError != nil {
	// 	log.Println("failed to indsert dummy posts", insertError)
	// }
}

func getDB() *gorm.DB {
	return db
}
