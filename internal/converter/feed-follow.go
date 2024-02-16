package converter

import (
	db "github.com/makarellav/bloagg/internal/db/sqlc"
	"github.com/makarellav/bloagg/internal/entity"
)

func ToFeedFollowFromDB(dbFeedFollow db.FeedFollow) *entity.FeedFollow {
	return &entity.FeedFollow{
		ID:        dbFeedFollow.ID.Bytes,
		UserID:    dbFeedFollow.UserID.Bytes,
		FeedID:    dbFeedFollow.FeedID.Bytes,
		CreatedAt: dbFeedFollow.CreatedAt.Time,
		UpdatedAt: dbFeedFollow.UpdatedAt.Time,
	}
}
