package newsfeed

import (
	"testing"
)

func TestAdd(t *testing.T) {
	feed := New()
	feed.Add(Item{})

	if len(feed.Items) == 0 {
		t.Error("Item was not added")
	}
}

func TestGetAll(t *testing.T) {
	feed := New()
	feed.Add(Item{})
	result := feed.GetAll()

	if len(feed.Items) != len(result) {
		t.Error("Didn't Get All Post")
	}
}
