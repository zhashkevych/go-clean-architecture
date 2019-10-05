package usecase

import (
	"context"
	"github.com/zhashkevych/go-clean-architecture/auth"
	"github.com/zhashkevych/go-clean-architecture/bookmark"
)

type BookmarkUseCase struct {
	bookmarkRepo bookmark.Repository
}

func NewBookmarkUseCase(bookmarkRepo bookmark.Repository) *BookmarkUseCase {
	return &BookmarkUseCase{
		bookmarkRepo: bookmarkRepo,
	}
}

func (b BookmarkUseCase) CreateBookmark(ctx context.Context, user *auth.User, todo *bookmark.Bookmark) error {
	return b.bookmarkRepo.CreateBookmark(ctx, user, todo)
}

func (b BookmarkUseCase) GetBookmarks(ctx context.Context, user *auth.User) ([]*bookmark.Bookmark, error) {
	return b.bookmarkRepo.GetBookmarks(ctx, user)
}

func (b BookmarkUseCase) GetBookmarkByID(ctx context.Context, user *auth.User, id int64) (*bookmark.Bookmark, error) {
	return b.GetBookmarkByID(ctx, user, id)
}

func (b BookmarkUseCase) DeleteBookmark(ctx context.Context, user *auth.User, id int64) error {
	return b.bookmarkRepo.DeleteBookmark(ctx, user, id)
}
