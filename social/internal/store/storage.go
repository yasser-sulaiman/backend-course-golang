package store

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

var (
	ErrNotFound          = errors.New("resource not found")
	ErrConflict          = errors.New("resource already exists")
	QueryTimeoutDuration = 5 * time.Second
)

type Storage struct {
	Posts interface {
		Create(context.Context, *Post) error
		GetByID(context.Context, int64) (*Post, error)
		Delete(context.Context, int64) error
		Update(context.Context, *Post) error
	}

	Users interface {
		Create(context.Context, *User) error
		GetByID(context.Context, int64) (*User, error)
	}

	Comments interface {
		Create(context.Context, *Comment) error
		GetByPostID(ctx context.Context, postID int64) ([]Comment, error)
	}

	Followers interface {
		Follow(ctx context.Context, FollowerID, UserId int64) error
		UnFollow(ctx context.Context, FollowerID, UserId int64) error
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Posts:    &PostStore{db: db},
		Users:    &UserStore{db: db},
		Comments: &CommentStore{db: db},
		Followers: &FollowerStore{db: db},
	}
}
