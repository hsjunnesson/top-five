package main

import (
	"time"
	"log"
	"strconv"
	"errors"
	"github.com/hsjunnesson/top-five/Godeps/_workspace/src/github.com/gin-gonic/gin"
	"github.com/hsjunnesson/top-five/Godeps/_workspace/src/gopkg.in/redis.v3"
)


// Types

type List struct {
	Id          string
	Title       string
	CreatedAt   time.Time
	ListItems   [5]string
}

type Lists []List


// Database functions

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
		if err := client.RPush("list_items:" + strId, val).Err(); err != nil {
			log.Printf("Couldn't RPush list_items:" + strId, err)
			return "", err
		}
	}
	
	if err := client.LPush("lists", strId).Err(); err != nil {
		log.Printf("Couldn't LPush lists " + strId, err)
		return "", err
	}
	
	return strId, nil
}

func IsListDeletedDB(client *redis.Client, id string) (bool, error) {
	isMember, err := client.SIsMember("deleted_lists", id).Result()
	if err != nil {
		log.Printf("Couldn't SISMember deleted_lists " + id)
		return false, err
	}

	return isMember, nil
}

// Fetches a List from the connected redis.
func GetListDB(client *redis.Client, id string) (*List, error) {
	// First check if list has been marked as deleted
	isDeleted, err := IsListDeletedDB(client, id)
	if err != nil {
		return nil, err
	}

	if isDeleted {
		return nil, nil
	}

	// Fetch parts from redis
	parts, err := client.HMGet("list:" + id, "id", "title", "created_at").Result()
	if err != nil {
		log.Printf("Could not HMGet list:" + id, err)
		return nil, err
	}

	listId, ok := parts[0].(string)
	if !ok {
		return nil, errors.New("Not ok")
	}

	title, ok := parts[1].(string)
	if !ok {
		return nil, errors.New("Not ok")
	}

	created_at, ok := parts[2].(string)
	if !ok {
		return nil, errors.New("Not ok")
	}

	created_at_time, err := time.Parse(time.RFC3339, created_at)
	if err != nil {
		log.Printf("Couldn't parse time: " + created_at, err)
		return nil, err
	}

	items, err := client.LRange("list_items:" + id, 0, 5).Result()
	if err != nil {
		log.Printf("Couldn't LRange list_items:" + id, err)
		return nil, err
	}

	if len(items) != 5 {
		log.Printf("List items not five elements")
		return nil, errors.New("List items not five elements")
	}

	list := List{
		Id: listId,
		Title: title,
		CreatedAt: created_at_time,
		ListItems: [5]string{items[0], items[1], items[2], items[3], items[4]},
	}

	return &list, nil
}

// Deletes a list from the connected redis.
// This doesn't actually delete it, just marks it as deleted.
func DeleteListDB(client *redis.Client, id string) error {
	if err := client.SAdd("deleted_lists", id).Err(); err != nil {
		log.Printf("Couldn't SAdd deleted_lists " + id, err)
		return err
	}

	return nil
}


// Route handlers

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

