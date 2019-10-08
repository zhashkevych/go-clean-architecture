package usecase

import (
	"context"
	"github.com/google/uuid"
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

func (b BookmarkUseCase) CreateBookmark(ctx context.Context, user *auth.User, url, title string) error {
	bm := &bookmark.Bookmark{
		URL:   url,
		Title: title,
	}

	return b.bookmarkRepo.CreateBookmark(ctx, user, bm)
}

func (b BookmarkUseCase) GetBookmarks(ctx context.Context, user *auth.User) ([]*bookmark.Bookmark, error) {
	return b.bookmarkRepo.GetBookmarks(ctx, user)
}

func (b BookmarkUseCase) DeleteBookmark(ctx context.Context, user *auth.User, id uuid.UUID) error {
	return b.bookmarkRepo.DeleteBookmark(ctx, user, id)
}
