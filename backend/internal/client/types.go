package client

import (
	"errors"
	"time"
)

type Client struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Whatsapp  string    `json:"whatsapp"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ClientCreateRequest struct {
	Name     string `json:"name"`
	Whatsapp string `json:"whatsapp"`
	Email    string `json:"email"`
}

func (c *ClientCreateRequest) Validate() error {
	if c.Name == "" {
		return errors.New("a client must have a name")
	}

	if c.Whatsapp == "" {
		return errors.New("a client must have a whatsapp")
	}

	if c.Email == "" {
		return errors.New("a client must have a email")
	}

	return nil
}
