package main

import (
	"testing"
)

func TestLists(t *testing.T) {
	list := List{Id: "", Title: "This is a title", Items: [5]string{"one", "two", "three", "four", "five"}, Timestamp: time.Now()}
	id, err := CreateList(redisClient, list)
	if err != nil {
		t.Fatal(err)
	}
}

