package post

import "time"

type Media struct {
	Type string `json:"content_type"` // "image" or "video"
	URL  string `json:"post_url"`
}
type Post struct {
	ID            int       `json:"id"`
	UserID        int       `json:"user_id"`
	ChannelID     *int      `json:"channel_id,omitempty"`
	Media         []Media   `json:"media"`
	LinkedVideoID *int      `json:"linked_video_id,omitempty"`
	Caption       string    `json:"caption"`
	Status        string    `json:"status"`
	Visibility    string    `json:"visibility"`
	ViewCount     int       `json:"view_count"`
	LikeCount     int       `json:"like_count"`
	ShareCount    int       `json:"share_count"`
	Location      string    `json:"location,omitempty"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type CreatePostInput struct {
	ChannelID     *int    `json:"channel_id,omitempty"`
	Caption       string  `json:"caption"`
	Media         []Media `json:"media_id"`
	LinkedVideoID *int    `json:"linked_video_id,omitempty"`
	Visibility    string  `json:"visibility"`
	Location      string  `json:"location,omitempty"`
}
