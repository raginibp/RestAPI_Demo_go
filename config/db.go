package configs

import (
	"RestAPI_Demo/controllers"
	"github.com/go-pg/pg"
	"log"
	"os"
)

func Connect() *pg.DB {
	opts := &pg.Options{
		User:     "postgres",
		Password: "123456",
		Addr:     "localhost:5432",
		Database: "new_books",
	}
	var db *pg.DB = pg.Connect(opts)
	if db == nil {
		log.Printf("Failed to connect")
		os.Exit(100)
	}
	log.Printf("Connected to db")
	controllers.CreateAlbumTable(db)
	controllers.InitiateDB(db)
	return db
}
