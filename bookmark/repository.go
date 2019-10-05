package bookmark

import (
	"context"
	"github.com/zhashkevych/go-clean-architecture/bookmark/model"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *model.User) error
	GetUser(ctx context.Context, id int64) (*model.User, error)
}

type BookmarkRepository interface {
	CreateBookmark(ctx context.Context, user *model.User, todo *model.Bookmark) error
	GetBookmarks(ctx context.Context, user *model.User) ([]*model.Bookmark, error)
	GetBookmarkByID(ctx context.Context, user *model.User, id int64) (*model.Bookmark, error)
	DeleteBookmark(ctx context.Context, user *model.User, id int64) error
}
