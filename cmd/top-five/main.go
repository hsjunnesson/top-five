package main

import (
	"log"
	"net/http"

	"github.com/hsjunnesson/top-five/Godeps/_workspace/src/github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	r.Run(":8080")
}
