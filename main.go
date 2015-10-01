package main

import (
	"log"
	"os"
	"time"
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
	
	app.Run(":" + port)

}

