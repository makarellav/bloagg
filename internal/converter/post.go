package converter

import (
	db "github.com/makarellav/bloagg/internal/db/sqlc"
	"github.com/makarellav/bloagg/internal/entity"
)

func ToPostFromDB(post db.Post) *entity.Post {
	return &entity.Post{
		ID:          post.ID.Bytes,
		Title:       post.Title,
		URL:         post.Url,
		Description: post.Description.String,
		PublishedAt: post.PublishedAt.Time,
		CreatedAt:   post.CreatedAt.Time,
		UpdatedAt:   post.UpdatedAt.Time,
		FeedID:      post.FeedID.Bytes,
	}
}
