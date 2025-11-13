package models

import "time"

type Channels struct {
	ID              int       `json:"id"`
	UserID          int       `json:"user_id"`
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	Verified        bool      `json:"verified"`
	SubscriberCount int       `json:"subscriber_count"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
