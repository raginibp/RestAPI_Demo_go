package routes

import (
	"RestAPI_Demo/controllers"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	router.GET("/", controllers.Welcome)
	router.GET("/albums", controllers.GetAlbums)
	router.POST("/albumsr", controllers.PostAlbum)
	router.GET("/albums/:id", controllers.GetAlbumByID)
	router.PUT("/albums/:id", controllers.PutAlbumByID)
	router.DELETE("/albums/:id", controllers.DeleteAlbumByID)
}
