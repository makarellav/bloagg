package entity

import (
	"github.com/google/uuid"
	"time"
)

type Post struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	URL         string    `json:"url"`
	Description string    `json:"description"`
	PublishedAt time.Time `json:"published_at"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	FeedID      uuid.UUID `json:"feed_id"`
}
