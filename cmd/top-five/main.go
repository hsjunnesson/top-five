package main

import (
	"log"
	"net/http"
	"os"

	"github.com/hsjunnesson/top-five/Godeps/_workspace/src/github.com/gin-gonic/gin"
)

func main() {
	var err error
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}
	
	router := gin.New()
	route.Use(gin.Logger())
	
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	
	r.Run(":" + port)
}

