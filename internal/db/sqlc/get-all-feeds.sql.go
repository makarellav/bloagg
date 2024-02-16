// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: get-all-feeds.sql

package db

import (
	"context"
)

const getAllFeeds = `-- name: GetAllFeeds :many
SELECT id, created_at, updated_at, name, url, user_id FROM feeds
`

func (q *Queries) GetAllFeeds(ctx context.Context) ([]Feed, error) {
	rows, err := q.db.Query(ctx, getAllFeeds)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Feed
	for rows.Next() {
		var i Feed
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Name,
			&i.Url,
			&i.UserID,
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