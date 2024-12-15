-- name: CreatePost :one
INSERT INTO posts (
    id,
    name,
    url,
    feed_id,
    created_at
) VALUES (
    $1, $2, $3, $4, $5
)
RETURNING *;

-- name: GetPosts