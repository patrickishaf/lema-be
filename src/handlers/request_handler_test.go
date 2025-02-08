package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/patrickishaf/lema-be/src/models"
	"github.com/stretchr/testify/assert"
)

func TestGetUsers(t *testing.T) {
	router := gin.Default()
	router.GET("/users", GetUsers)

	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Fatalf("failed to send request %v", err)
	}

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code, "Response does not have status 200")

	var users []models.User
	err = json.Unmarshal(recorder.Body.Bytes(), &users)
	if err != nil {
		t.Fatalf("failed to parse response body. error: %v", err)
	}
}
