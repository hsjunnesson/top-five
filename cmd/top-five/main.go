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
	
	router := gin.New()
	router.Use(gin.Logger())
	
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	
	router.Run(":" + port)
}

