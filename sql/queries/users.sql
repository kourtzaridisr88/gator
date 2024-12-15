-- name: CreateUser :one
INSERT INTO users (id, name, created_at, updated_at)
VALUES (
    $1,
    $2,
    $3,
    $4
)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users WHERE id = $1 LIMIT 1;

-- name: GetUserByName :one
SELECT * FROM users WHERE name = $1 LIMIT 1;

-- name: TruncateUsers :exec
TRUNCATE users CASCADE;

-- name: ListUsers :many
SELECT name FROM users;