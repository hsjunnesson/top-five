package main

import (
	"time"
	"log"
	"strconv"
	"github.com/hsjunnesson/top-five/Godeps/_workspace/src/github.com/gin-gonic/gin"
	"github.com/hsjunnesson/top-five/Godeps/_workspace/src/gopkg.in/redis.v3"
)

type List struct {
	Id          string
	Title       string
	CreatedAt   time.Time
	ListItems   [5]string
}

type Lists []List

// Creates a List in the connected redis.
// Returns the id of the List, or error.
func CreateListDB(client *redis.Client, title string, listItems [5]string) (string, error) {
	timestamp := time.Now().Format(time.RFC3339)
	
	id, err := client.Incr("next_list_id").Result()
	if err != nil {
		log.Printf("Couldn't increment next_list_id", err)
		return "", err
	}
	
	strId := strconv.FormatInt(id, 10)
	
	if client.HMSet(
		"list:" + strId,
		"id", strId,
		"title", title,
		"created_at", timestamp,		
	).Err() != nil {
		log.Printf("Couldn't HMSet list:" + strId, err)
		return "", err
	}
	
	for _, val := range listItems {
		if client.RPush("list_items:" + strId, val).Err() != nil {
			log.Printf("Couldn't RPush list_items:" + strId, err)
			return "", err
		}
	}
	
	if client.LPush("lists", strId).Err() != nil {
		log.Printf("Couldn't LPush lists " + strId, err)
		return "", err
	}
	
	return strId, nil
}

// Fetches a List from the connected redis.
func GetListDB(client *redis.Client, id string) (*List, error) {
	parts, err := client.HMGet("list:" + id, "id", "title", "created_at").Result()
	if err != nil {
		log.Printf("Could not HMGet list:" + id, err)
		return nil, err
	}

	id, err := parts[0].(string)
	if err != nil {
		return nil, err
	}

	title, err := parts[1].(string)
	if err != nil {
		return nil, err
	}

		
	
	var items = [5]string{"one", "two", "three", "four", "five"}
	list := List{
		Id: parts[0],
		Title: parts[1],
		CreatedAt: time.New(parts[2]),
		ListItems: items,
	}

	return list, nil
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

