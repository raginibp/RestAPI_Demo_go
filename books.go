package main

import (
	"database/sql"
	"fmt"
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "123456"
	DB_NAME     = "books"
)

func main() {
	psqlconn := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", psqlconn)
	defer db.Close()
	print(err)

}

//func setupDB() *sql.DB {
//	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
//	db, err := sql.Open("postgres", dbinfo)
//
//	return db
//}
