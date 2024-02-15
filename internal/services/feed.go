package services

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/makarellav/bloagg/internal/converter"
	db "github.com/makarellav/bloagg/internal/db/sqlc"
	"github.com/makarellav/bloagg/internal/entity"
	"time"
)

type FeedService struct {
	db  *db.Queries
	ctx context.Context
}

func NewFeedService(db *db.Queries, ctx context.Context) *FeedService {
	return &FeedService{db: db, ctx: ctx}
}

func (fs *FeedService) Create(feed *entity.Feed) (*entity.Feed, error) {
	feedParams := db.CreateFeedParams{
		ID: pgtype.UUID{
			Bytes: uuid.New(),
			Valid: true,
		},
		Name: feed.Name,
		Url:  feed.URL,
		CreatedAt: pgtype.Timestamp{
			Time:  time.Now(),
			Valid: true,
		},
		UpdatedAt: pgtype.Timestamp{
			Time:  time.Now(),
			Valid: true,
		},
		UserID: pgtype.UUID{
			Bytes: feed.UserID,
			Valid: true,
		},
	}

	dbFeed, err := fs.db.CreateFeed(fs.ctx, feedParams)

	if err != nil {
		return nil, err
	}

	newFeed := converter.ToFeedFromDB(dbFeed)

	return newFeed, nil
}
