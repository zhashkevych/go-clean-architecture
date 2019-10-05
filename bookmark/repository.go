package bookmark

import (
	"context"
	"github.com/google/uuid"
	"github.com/zhashkevych/go-clean-architecture/auth"
)

type Repository interface {
	CreateBookmark(ctx context.Context, user *auth.User, todo *Bookmark) error
	GetBookmarks(ctx context.Context, user *auth.User) ([]*Bookmark, error)
	GetBookmarkByID(ctx context.Context, user *auth.User, id uuid.UUID) (*Bookmark, error)
	DeleteBookmark(ctx context.Context, user *auth.User, id uuid.UUID) error
}
