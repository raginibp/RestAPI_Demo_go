package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"log"
	"net/http"
	"time"
)

type album struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Artist    string    `json:"artist"`
	Price     float64   `json:"price"`
	Completed string    `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

//var albums = []album{
//	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
//	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
//	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
//}

func CreateAlbumTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createError := db.CreateTable(&album{}, opts)

	if createError != nil {
		log.Printf("Error while creating Album table, Reason: %v\n", createError)
		return createError
	}
	log.Printf("album table created")
	return nil
}

var dbConnect *pg.DB

func InitiateDB(db *pg.DB) {
	dbConnect = db
}

func GetAlbums(c *gin.Context) {
	var Album []album
	err := dbConnect.Model(&Album).Select()
	if err != nil {
		log.Printf("Error while getting all todos, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "All Albums",
		"data":    Album,
	})
	return
}

func PostAlbum(c *gin.Context) {
	var newAlbum album

	c.BindJSON(&newAlbum)
	id := newAlbum.ID
	title := newAlbum.Title
	artist := newAlbum.Artist
	price := newAlbum.Price
	completed := newAlbum.Completed
	createdAt := newAlbum.CreatedAt
	updatedAt := newAlbum.UpdatedAt

	insertError := dbConnect.Insert(&album{
		ID:        id,
		Title:     title,
		Artist:    artist,
		Price:     price,
		Completed: completed,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	})
	if insertError != nil {
		log.Printf("Error while inserting new album into db, Reason: %v\n", insertError)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Album created Successfully",
	})
	return
}

func GetAlbumByID(c *gin.Context) {
	id := c.Param("id")
	album := &album{ID: id}
	err := dbConnect.Select(album)

	if err != nil {
		log.Printf("Error while getting a single todo, Reason: %v\n", err)
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Album not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Single album",
		"data":    album,
	})
	return
}

func PutAlbumByID(c *gin.Context) {
	id := c.Param("id")
	var Album album
	c.BindJSON(&Album)
	completed := Album.Completed

	_, err := dbConnect.Model(&album{}).Set("completed = ?", completed).Where("id =?", id).Update()
	if err != nil {
		log.Printf("Error, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "Something went wrong",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Album Edited successfully",
	})
	return
}

func DeleteAlbumByID(c *gin.Context) {
	id := c.Param("id")
	Album := &album{ID: id}
	err := dbConnect.Delete(Album)
	if err != nil {
		log.Printf("Error while deleting a single album, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Album deleted successfully",
	})
}
