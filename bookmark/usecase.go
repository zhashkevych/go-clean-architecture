package bookmark

import (
	"context"
	"github.com/zhashkevych/go-clean-architecture/auth"
)

type UseCase interface {
	CreateBookmark(ctx context.Context, user *auth.User, todo *Bookmark) error
	GetBookmarks(ctx context.Context, user *auth.User) ([]*Bookmark, error)
	GetBookmarkByID(ctx context.Context, user *auth.User, id int64) (*Bookmark, error)
	DeleteBookmark(ctx context.Context, user *auth.User, id int64) error
}
