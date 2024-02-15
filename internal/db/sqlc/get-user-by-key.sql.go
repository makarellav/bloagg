// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: get-user-by-key.sql

package db

import (
	"context"
)

const getUserByKey = `-- name: GetUserByKey :one
SELECT id, created_at, updated_at, name, api_key FROM users
WHERE api_key = $1
`

func (q *Queries) GetUserByKey(ctx context.Context, apiKey string) (User, error) {
	row := q.db.QueryRow(ctx, getUserByKey, apiKey)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.ApiKey,
	)
	return i, err
}
