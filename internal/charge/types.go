package charge

import (
	"errors"
	"time"
)

type Charge struct {
	ID               int       `json:"id"`
	Value            int       `json:"value"`
	Observation      string    `json:"observation"`
	NotificationDate time.Time `json:"notification_date"`
	ClientId         int       `json:"client_id"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

type CreateChargeRequest struct {
	Value            int       `json:"value"`
	Observation      string    `json:"observation"`
	NotificationDate time.Time `json:"notification_date"`
	ClientId         int       `json:"client_id"`
}

func (c *CreateChargeRequest) Validate() error {
	if c.Value == 0 {
		return errors.New("charge value must not be 0")
	}

	if c.Observation == "" {
		return errors.New("a charge must have a observation")
	}

	if c.NotificationDate.IsZero() {
		return errors.New("a charge must have a notification date")
	}

	if c.ClientId == 0 {
		return errors.New("a charge must have a client")
	}

	return nil
}
