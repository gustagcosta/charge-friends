package main

import "time"

type Charge struct {
	ID               int       `json:"id"`
	Value            int       `json:"value"`
	Observation      string    `json:"observation"`
	NotificationDate time.Time `json:"notification_date"`
	ClientId         int       `json:"client_id"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}
