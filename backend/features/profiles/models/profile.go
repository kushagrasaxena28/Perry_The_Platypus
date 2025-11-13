package models

import "time"

type Profile struct {
	ID             int       `json:"id"`
	Username       string    `json:"username"`
	Email          string    `json:"email"`
	PasswordHash   string    `json:"-"`
	Privacy        string    `json:"privacy"`
	Verified       bool      `json:"verified"`
	FollowerCount  int       `json:"follower_count"`
	FollowingCount int       `json:"following_count"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
