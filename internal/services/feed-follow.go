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

type FeedFollowService struct {
	db  *db.Queries
	ctx context.Context
}

func NewFeedFollowService(db *db.Queries, ctx context.Context) *FeedFollowService {
	return &FeedFollowService{
		db:  db,
		ctx: ctx,
	}
}

func (ffs *FeedFollowService) Create(follow *entity.FeedFollow) (*entity.FeedFollow, error) {
	dbFeedFollowParams := db.CreateFeedFollowParams{
		ID: pgtype.UUID{
			Bytes: uuid.New(),
			Valid: true,
		},
		UserID: pgtype.UUID{
			Bytes: follow.UserID,
			Valid: true,
		},
		FeedID: pgtype.UUID{
			Bytes: follow.FeedID,
			Valid: true,
		},
		CreatedAt: pgtype.Timestamp{
			Time:  time.Now(),
			Valid: true,
		},
		UpdatedAt: pgtype.Timestamp{
			Time:  time.Now(),
			Valid: true,
		},
	}
	dbFeedFollow, err := ffs.db.CreateFeedFollow(ffs.ctx, dbFeedFollowParams)

	if err != nil {
		return nil, err
	}

	return converter.ToFeedFollowFromDB(dbFeedFollow), nil
}

func (ffs *FeedFollowService) Delete(follow entity.FeedFollow) error {
	deleteParams := db.DeleteFeedFollowParams{
		UserID: pgtype.UUID{
			Bytes: follow.UserID,
			Valid: true,
		},
		ID: pgtype.UUID{
			Bytes: follow.ID,
			Valid: true,
		},
	}

	return ffs.db.DeleteFeedFollow(ffs.ctx, deleteParams)
}

func (ffs *FeedFollowService) GetAllByUserID(userID uuid.UUID) ([]*entity.FeedFollow, error) {
	dbFeedFollows, err := ffs.db.GetFeedFollowsByUserID(ffs.ctx, pgtype.UUID{
		Bytes: userID,
		Valid: true,
	})

	if err != nil {
		return nil, err
	}

	feedFollows := make([]*entity.FeedFollow, len(dbFeedFollows))

	for i, feedFollow := range dbFeedFollows {
		feedFollows[i] = converter.ToFeedFollowFromDB(feedFollow)
	}

	return feedFollows, nil
}
