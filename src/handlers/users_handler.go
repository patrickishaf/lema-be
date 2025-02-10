package handlers

import (
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/patrickishaf/lema-be/src/common"
	"github.com/patrickishaf/lema-be/src/db"
)

func GetUsers(c *gin.Context) {
	pageNumber, _ := strconv.Atoi(c.DefaultQuery("pageNumber", "0"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	offset := pageNumber * pageSize
	users := db.FindUsers(pageSize, offset)
	count := db.FindUserCount()
	totalPages := math.Ceil(float64(count) / float64(pageSize))

	c.IndentedJSON(http.StatusOK, map[string]any{
		"pageNumber": pageNumber,
		"pageSize":   pageSize,
		"totalPages": totalPages,
		"data":       users,
	})
}

func GetUserCount(c *gin.Context) {
	numberOfusers := db.FindUserCount()
	c.IndentedJSON(http.StatusOK, numberOfusers)
}

func GetUserById(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, common.CreateErrorResponse("invalid id param"))
		return
	}

	existingUser := db.FindUserById(uint(userID))
	c.IndentedJSON(http.StatusOK, existingUser)
}
