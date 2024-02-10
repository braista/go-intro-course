-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetFeeds :many
SELECT * 
FROM feeds;

-- name: GetUserFeeds :many
SELECT f.id, f.created_at, f.updated_at, f.name, f.url, f.user_id
FROM feeds f
LEFT JOIN feeds_follows ff ON f.id = ff.feed_id
WHERE ff.user_id = $1;

-- name: CreateFeedFollow :one
INSERT INTO feeds_follows (id, feed_id, user_id, created_at)
VALUES ($1, $2, $3, $4)
RETURNING *;