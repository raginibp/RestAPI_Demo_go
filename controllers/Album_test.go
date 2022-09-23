package controllers

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router

}
func TestPostAlbums(t *testing.T) {
	r := SetUpRouter()
	r.POST("/albumsr", PostAlbum)
	Album := album{
		ID:        "78",
		Title:     "abc",
		Artist:    "abc",
		Price:     5,
		Completed: "d",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	jsonValue, _ := json.Marshal(Album)
	req, err := http.NewRequest("POST", "/albumsr", bytes.NewBuffer(jsonValue))
	if err != nil {
		t.Fatal(err)
	}
	W := httptest.NewRecorder()
	//fmt.Println(req)
	r.ServeHTTP(W, req)
	assert.Equal(t, http.StatusCreated, W.Code)
}

func TestGetAlbums(t *testing.T) {
	engine := Engine()
}
