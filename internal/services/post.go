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

type PostService struct {
	db  *db.Queries
	ctx context.Context
}

func NewPostService(db *db.Queries, ctx context.Context) *PostService {
	return &PostService{
		db:  db,
		ctx: ctx,
	}
}

func (ps *PostService) Create(item *entity.RSSFeed, feedID uuid.UUID) error {
	var publishedAt pgtype.Timestamp

	if t, err := time.Parse(time.RFC1123, item.PubDate); err == nil {
		publishedAt = pgtype.Timestamp{
			Time:  t,
			Valid: true,
		}
	} else {
		t, err = time.Parse(time.RFC1123Z, item.PubDate)

		if err == nil {
			publishedAt = pgtype.Timestamp{
				Time:  t,
				Valid: true,
			}
		}
	}

	createPostParams := db.CreatePostParams{
		ID: pgtype.UUID{
			Bytes: uuid.New(),
			Valid: true,
		},
		Title: item.Title,
		Url:   item.Link,
		Description: pgtype.Text{
			String: item.Description,
			Valid:  true,
		},
		PublishedAt: publishedAt,
		CreatedAt: pgtype.Timestamp{
			Time:  time.Now(),
			Valid: true,
		},
		UpdatedAt: pgtype.Timestamp{
			Time:  time.Now(),
			Valid: true,
		},
		FeedID: pgtype.UUID{
			Bytes: feedID,
			Valid: true,
		},
	}

	return ps.db.CreatePost(ps.ctx, createPostParams)
}

func (ps *PostService) GetByUserID(userID uuid.UUID, limit int) ([]*entity.Post, error) {
	dbPosts, err := ps.db.GetPostsByUserID(ps.ctx, db.GetPostsByUserIDParams{
		UserID: pgtype.UUID{
			Bytes: userID,
			Valid: true,
		},
		Limit: int32(limit),
	})

	if err != nil {
		return nil, err
	}

	userPosts := make([]*entity.Post, len(dbPosts))

	for i, dbPost := range dbPosts {
		userPosts[i] = converter.ToPostFromDB(dbPost)
	}

	return userPosts, nil
}
