package main

import (
	"RestAPI_Demo/config"
	"RestAPI_Demo/routes"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	configs.Connect()
	router := gin.Default()
	routes.Routes(router)
	log.Fatal(router.Run(":3000"))
}
