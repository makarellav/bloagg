-- name: GetFeedFollowsByUserID :many
SELECT * FROM feed_follows WHERE user_id = $1;