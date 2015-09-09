package hn

import (
	"fmt"
	"time"
)

var errMissingID = fmt.Errorf("missing id")

// UsersService communicates with the news
// related endpoints in the Hacker News API
type UsersService interface {
	Get(id string) (*User, error)
}

// usersService implements LiveService.
type usersService struct {
	client *Client
}

// User represents a Hacker News user
type User struct {
	About     string `json:"about"`
	Created   int    `json:"created"`
	Delay     int    `json:"delay"`
	ID        string `json:"id"`
	Karma     int    `json:"karma"`
	Submitted []int  `json:"submitted"`
}

// CreatedTime return the time of the created
func (u *User) CreatedTime() time.Time {
	return time.Unix(int64(u.Created), 0)
}

// User is a convenience method proxying Users.Get
func (c *Client) User(id string) (*User, error) {
	return c.Users.Get(id)
}

// Get retrieves a user with the given id
func (s *usersService) Get(id string) (*User, error) {
	if id == "" {
		return nil, errMissingID
	}

	req, err := s.client.NewRequest(s.getPath(id))
	if err != nil {
		return nil, err
	}

	var user User
	_, err = s.client.Do(req, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *usersService) getPath(id string) string {
	return fmt.Sprintf("user/%v.json", id)
}
