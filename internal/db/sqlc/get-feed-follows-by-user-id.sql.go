// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: get-feed-follows-by-user-id.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const getFeedFollowsByUserID = `-- name: GetFeedFollowsByUserID :many
SELECT id, user_id, feed_id, created_at, updated_at FROM feed_follows WHERE user_id = $1
`

func (q *Queries) GetFeedFollowsByUserID(ctx context.Context, userID pgtype.UUID) ([]FeedFollow, error) {
	rows, err := q.db.Query(ctx, getFeedFollowsByUserID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FeedFollow
	for rows.Next() {
		var i FeedFollow
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.FeedID,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
