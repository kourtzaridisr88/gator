-- name: CreateFeed :one
INSERT INTO feeds (id, name, url, user_id)
VALUES (
    $1,
    $2,
    $3,
    $4
)
RETURNING *;

-- name: ListFeeds :many
SELECT f.name, f.url, u.name AS username FROM feeds f
INNER JOIN users u ON u.id = f.user_id;

-- name: GetFeedByUrl :one
SELECT * FROM feeds
WHERE url = $1;

-- name: ListFeedsByUser :many
SELECT f.name FROM feeds f
INNER JOIN feed_follows ff ON ff.feed_id = f.id
WHERE ff.user_id = $1;

-- name: MarkFeedFetched :exec
UPDATE feeds SET last_fetched_at = NOW(), updated_at = NOW()
WHERE id = $1;

-- name: GetNextFeedToFetch :one
SELECT * FROM feeds 
ORDER BY last_fetched_at NULLS FIRST
LIMIT 1;