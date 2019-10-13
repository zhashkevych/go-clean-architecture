package bookmark

import (
	"context"
	"github.com/zhashkevych/go-clean-architecture/auth"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Repository interface {
	CreateBookmark(ctx context.Context, user *auth.User, bm *Bookmark) error
	GetBookmarks(ctx context.Context, user *auth.User) ([]*Bookmark, error)
	DeleteBookmark(ctx context.Context, user *auth.User, id primitive.ObjectID) error
}
