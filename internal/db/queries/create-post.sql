-- name: CreatePost :exec
INSERT INTO posts(id, title, url, description, published_at, created_at, updated_at, feed_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8);