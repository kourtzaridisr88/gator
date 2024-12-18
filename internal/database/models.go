// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package database

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Feed struct {
	ID            uuid.UUID
	Name          string
	Url           string
	UserID        uuid.UUID
	CreatedAt     time.Time
	UpdatedAt     sql.NullTime
	LastFetchedAt sql.NullTime
}

type FeedFollow struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	FeedID    uuid.UUID
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}

type User struct {
	ID        uuid.UUID
	Name      string
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}
