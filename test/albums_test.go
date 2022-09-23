package t_test

import (
	"RestAPI_Demo/controllers"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router

}
func TestWelcome(t *testing.T) {
	mockResponse := `{"message":"Welcome to Go Rest Api","status":200}`

	r := SetUpRouter()
	r.GET("/", controllers.Welcome)
	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}

//func TestGetAlbums(t *testing.T) {
//
//	expected := ` `
//	r := SetUpRouter()
//	r.GET("/albums", controllers.GetAlbums)
//	req, _ := http.NewRequest("GET", "/albums", nil)
//	w := httptest.NewRecorder()
//	if status := w.Code; status != http.StatusOK {
//		t.Errorf("handler returned wrong status code: fot %v want %v", status, http.StatusOK)
//	}
//	if w.Body.String() != expected {
//		t.Errorf("handler returned unexpected body: got %v want %v", w.Body.String(), expected)
//	}
//	r.ServeHTTP(w, req)
//}

//func TestPostAlbums(t *testing.T) {
//	r := SetUpRouter()
//	r.POST("/albums", controllers.PostAlbum)
//	albumId :=
//	album := album{
//		I
//	}
//}
