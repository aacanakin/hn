package hn

import (
	"fmt"
	"time"
)

// ItemsService communicates with the news
// related endpoints in the Hacker News API
type ItemsService interface {
	Get(id int) (*Item, error)
}

// itemsService implements ItemsService.
type itemsService struct {
	client *Client
}

// Item represents a item
type Item struct {
	ID        int    `json:"id" bson:"_id"`
	Parent    int    `json:"parent" bson:"parent"`
	Kids      []int  `json:"kids" bson:"kids"`
	Parts     []int  `json:"parts" bson:"parts"`
	Score     int    `json:"score" bson:"score"`
	Timestamp int    `json:"time" bson:"time"`
	By        string `json:"by" bson:"by"`
	Type      string `json:"type" bson:"type"`
	Title     string `json:"title" bson:"title"`
	Text      string `json:"text" bson:"text"`
	URL       string `json:"url" bson:"url"`
	Dead      bool   `json:"dead" bson:"dead"`
	Deleted   bool   `json:"deleted" bson:"deleted"`
}

// Time return the time of the timestamp
func (i *Item) Time() time.Time {
	return time.Unix(int64(i.Timestamp), 0)
}

// Item is a convenience method proxying Items.Get
func (c *Client) Item(id int) (*Item, error) {
	return c.Items.Get(id)
}

// Get retrieves an item with the given id
func (s *itemsService) Get(id int) (*Item, error) {
	req, err := s.client.NewRequest(s.getPath(id))
	if err != nil {
		return nil, err
	}

	var item Item
	_, err = s.client.Do(req, &item)
	if err != nil {
		return nil, err
	}

	if item.Type == "story" && item.URL == "" {
		item.URL = fmt.Sprintf("https://news.ycombinator.com/item?id=%v", id)
	}

	return &item, nil
}

func (s *itemsService) getPath(id int) string {
	return fmt.Sprintf("item/%v.json", id)
}
