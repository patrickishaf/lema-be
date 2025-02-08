package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/patrickishaf/lema-be/src/db"
)

func GetUsers(c *gin.Context) {
	pageNumber, _ := strconv.Atoi(c.DefaultQuery("pageNumber", "0"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	offset := pageNumber * pageSize
	users := db.FindUsers(pageSize, offset)

	c.IndentedJSON(http.StatusOK, map[string]any{
		"pageNumber": pageNumber,
		"pageSize":   pageSize,
		"data":       users,
	})
}

func GetUserCount(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "get user count",
	})
}

func GetUserById(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "get user by id",
	})
}
