package entity

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	ID        uuid.UUID `json:"id,omitempty"`
	ApiKey    string    `json:"api_key,omitempty"`
}
