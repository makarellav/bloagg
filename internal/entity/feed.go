package entity

import (
	"github.com/google/uuid"
	"time"
)

type Feed struct {
	ID        uuid.UUID `json:"id,omitempty,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	Name      string    `json:"name"`
	URL       string    `json:"url"`
	UserID    uuid.UUID `json:"user_id,omitempty"`
}
