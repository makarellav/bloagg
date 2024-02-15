package converter

import (
	db "github.com/makarellav/bloagg/internal/db/sqlc"
	"github.com/makarellav/bloagg/internal/entity"
)

func ToFeedFromDB(dbFeed db.Feed) *entity.Feed {
	return &entity.Feed{
		ID:        dbFeed.ID.Bytes,
		CreatedAt: dbFeed.CreatedAt.Time,
		UpdatedAt: dbFeed.UpdatedAt.Time,
		Name:      dbFeed.Name,
		URL:       dbFeed.Url,
		UserID:    dbFeed.UserID.Bytes,
	}
}
