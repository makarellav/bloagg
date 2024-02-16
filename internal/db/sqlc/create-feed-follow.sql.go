// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: create-feed-follow.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createFeedFollow = `-- name: CreateFeedFollow :one
INSERT INTO feed_follows(id, user_id, feed_id, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, user_id, feed_id, created_at, updated_at
`

type CreateFeedFollowParams struct {
	ID        pgtype.UUID
	UserID    pgtype.UUID
	FeedID    pgtype.UUID
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
}

func (q *Queries) CreateFeedFollow(ctx context.Context, arg CreateFeedFollowParams) (FeedFollow, error) {
	row := q.db.QueryRow(ctx, createFeedFollow,
		arg.ID,
		arg.UserID,
		arg.FeedID,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i FeedFollow
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.FeedID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
