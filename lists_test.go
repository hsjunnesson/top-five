package main

import (
	"testing"
)

func TestLists(t *testing.T) {
	listId, err := CreateListDB(redisClient, "This is a title", [5]string{"one", "two", "three", "four", "five"})
	if err != nil {
		t.Fatal(err)
	}

  _, err := GetListDB(listId)
	if err != nil {
		t.Fatal(err)
	}

	
}

