package main

import (
	"errors"
	"strings"
	"testing"
)

func TestListsCRUD(t *testing.T) {
	title := "Top Five Primes"
	listItems := [5]string{"two", "three", "five", "seven", "eleven"}
	
	listId, err := CreateListDB(redisClient, title, listItems)
	if err != nil {
		t.Fatal(err)
	}

  list, err := GetListDB(redisClient, listId)
	if err != nil {
		t.Fatal(err)
	}
	
	if list.Id != listId {
		t.Fatal(errors.New("List id " + listId + " doesn't match id " + list.Id))
	}

	if list.Title != title {
		t.Fatal(errors.New("List title " + list.Title + " doesn't match title " + title))
	}

	if list.ListItems != listItems {
		t.Fatal(errors.New("List items '" + strings.Join(list.ListItems[:], ", ") + "' doesn't match list items '" + strings.Join(listItems[:], ", ") + "'"))
	}

	deleteErr := DeleteListDB(redisClient, listId)
	if deleteErr != nil {
		t.Fatal(deleteErr)
	}
	
  deletedList, err := GetListDB(redisClient, listId)
	if err != nil {
		t.Fatal(err)
	}
	if deletedList != nil {
		t.Fatal(errors.New("Can get a deleted list"))
	}
}

