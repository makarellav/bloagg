package entity

import (
	"github.com/google/uuid"
	"time"
)

type FeedFollow struct {
	ID        uuid.UUID `json:"id,omitempty"`
	UserID    uuid.UUID `json:"user_id,omitempty"`
	FeedID    uuid.UUID `json:"feed_id"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
