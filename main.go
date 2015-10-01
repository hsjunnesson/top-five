package main

import (
	"log"
	"os"
	"github.com/hsjunnesson/top-five/Godeps/_workspace/src/github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	app := gin.Default()
	app.GET("/lists", ListsList)
	app.GET("/lists/:list_id", ListsDetail)

	app.Static("/public", "./public")
	app.Static("/1726D19535E7EF78AD05045709BA5EA1.txt", "./public/1726D19535E7EF78AD05045709BA5EA1.txt")
	
	app.Run(":" + port)

}

