// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: feed_follows.sql

package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const createFeedFollow = `-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
    INSERT INTO feed_follows (id, user_id, feed_id)
    VALUES (
        $1,
        $2,
        $3
    )
    RETURNING id, user_id, feed_id, created_at, updated_at
)

SELECT
    inserted_feed_follow.id, inserted_feed_follow.user_id, inserted_feed_follow.feed_id, inserted_feed_follow.created_at, inserted_feed_follow.updated_at,
    feeds.name AS feed_name,
    users.name AS user_name
FROM inserted_feed_follow
INNER JOIN users ON users.id = inserted_feed_follow.user_id
INNER JOIN feeds ON feeds.id = inserted_feed_follow.feed_id
`

type CreateFeedFollowParams struct {
	ID     uuid.UUID
	UserID uuid.UUID
	FeedID uuid.UUID
}

type CreateFeedFollowRow struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	FeedID    uuid.UUID
	CreatedAt time.Time
	UpdatedAt sql.NullTime
	FeedName  string
	UserName  string
}

func (q *Queries) CreateFeedFollow(ctx context.Context, arg CreateFeedFollowParams) (CreateFeedFollowRow, error) {
	row := q.db.QueryRowContext(ctx, createFeedFollow, arg.ID, arg.UserID, arg.FeedID)
	var i CreateFeedFollowRow
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.FeedID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.FeedName,
		&i.UserName,
	)
	return i, err
}

const deleteFeedFollowByUserAndFeed = `-- name: DeleteFeedFollowByUserAndFeed :exec
DELETE FROM feed_follows
WHERE user_id = $1
AND feed_id = $2
`

type DeleteFeedFollowByUserAndFeedParams struct {
	UserID uuid.UUID
	FeedID uuid.UUID
}

func (q *Queries) DeleteFeedFollowByUserAndFeed(ctx context.Context, arg DeleteFeedFollowByUserAndFeedParams) error {
	_, err := q.db.ExecContext(ctx, deleteFeedFollowByUserAndFeed, arg.UserID, arg.FeedID)
	return err
}
