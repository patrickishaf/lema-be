package main

import (
	"github.com/gin-gonic/gin"
	"github.com/patrickishaf/lema-be/src/db"
	"github.com/patrickishaf/lema-be/src/handlers"
)

func main() {
	db.InitializeDb()

	router := gin.Default()
	router.GET("/users", handlers.GetUsers)
	router.GET("/users/count", handlers.GetUserCount)
	router.GET("/users/:id", handlers.GetUserById)
	handlers.RegisterPostHandlers(router)

	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
