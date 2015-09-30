package main

import (
	"time"
	"log"
	"strconv"
	"github.com/hsjunnesson/top-five/Godeps/_workspace/src/github.com/gin-gonic/gin"
	"github.com/hsjunnesson/top-five/Godeps/_workspace/src/gopkg.in/redis.v3"
)

type List struct {
	Id          string    `json:"id"`
	Title       string    `json:"title"`
	Items       [5]string `json:"items"`
	Timestamp   time.Time `json:"timestamp"`
}

type Lists []List

func CreateList(client *redis.Client, l List) (int64, error) {
	timestamp := l.Timestamp.Format(time.RFC3339)
	
	id, err := client.Incr("next_list_id").Result()
	if err != nil {
		log.Printf("Couldn't increment next_list_id", err)
		return 0, err
	}
	
	strId := strconv.FormatInt(id, 10)
	
	if client.HMSet(
		"list:" + strId,
		"Id", strId,
		"Title", l.Title,
		"Timestamp", timestamp,		
	).Err() != nil {
		log.Printf("Couldn't HMSet list:" + strId, err)
		return 0, err
	}
	
	for _, val := range l.Items {
		if client.RPush("list_items:" + strId, val).Err() != nil {
			log.Printf("Couldn't RPush list_items:" + strId, err)
			return 0, err
		}
	}
	
	if client.LPush("lists", strId).Err() != nil {
		log.Printf("Couldn't LPush lists " + strId, err)
		return 0, err
	}
	
	return id, nil
}

func ListsList(c *gin.Context) {
	c.JSON(200, "Hej")
}

func ListsDetail(c *gin.Context) {
//	list_id := c.Params.ByName("id")
//	_, _ := strconv.Atoi(list_id)
//	list := lists()[0]
	//content := gin.H{"title": list.Title, "content": list.Items}
	c.JSON(200, "Hoj")
}

