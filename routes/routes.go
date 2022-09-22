package routes

import (
	"RestAPI_Demo/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Routes(router *gin.Engine) {
	router.GET("/", Welcome)
	router.GET("/albums", controllers.GetAlbums)
	router.POST("/albums", controllers.PostAlbum)
	router.GET("/albums/:id", controllers.GetAlbumByID)
	router.PUT("/albums/:id", controllers.PutAlbumByID)
	router.DELETE("/albums/:id", controllers.DeleteAlbumByID)
}
func Welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Welcome to Go Rest Api",
	})
	return
}
