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

	list := List{Id: "", Title: "This is a title", Items: [5]string{"one", "two", "three", "four", "five"}, Timestamp: time.Now()}
	id, err := CreateList(redisClient, list)
	if err != nil {
		log.Printf("error", err)
	} else {
		log.Printf("Id: %i", id)
	}	


	
	app := gin.Default()
	app.GET("/lists", ListsList)
	app.GET("/lists/:list_id", ListsDetail)
	
	app.Run(":" + port)

}

