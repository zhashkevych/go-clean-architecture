package usecase

import (
	"context"
	"github.com/zhashkevych/go-clean-architecture/bookmark"
	"github.com/zhashkevych/go-clean-architecture/bookmark/model"
)

type BookmarkUseCase struct {
	userRepo     bookmark.UserRepository
	bookmarkRepo bookmark.BookmarkRepository
}

func NewBookmarkUseCase(userRepo bookmark.UserRepository, bookmarkRepo bookmark.BookmarkRepository) *BookmarkUseCase {
	return &BookmarkUseCase{
		userRepo:     userRepo,
		bookmarkRepo: bookmarkRepo,
	}
}

func (b BookmarkUseCase) CreateUser(ctx context.Context, user *model.User) error {
	return b.userRepo.CreateUser(ctx, user)
}

func (b BookmarkUseCase) GetUser(ctx context.Context, id int64) (*model.User, error) {
	return b.userRepo.GetUser(ctx, id)
}

func (b BookmarkUseCase) CreateBookmark(ctx context.Context, user *model.User, todo *model.Bookmark) error {
	todo.UserID = user.ID

	return b.bookmarkRepo.CreateBookmark(ctx, todo)
}

func (b BookmarkUseCase) GetBookmarks(ctx context.Context, user *model.User) ([]*model.Bookmark, error) {
	return b.bookmarkRepo.GetBookmarksByUserID(ctx, user.ID)
}

func (b BookmarkUseCase) GetBookmarkByID(ctx context.Context, user *model.User, id int64) (*model.Bookmark, error) {
	return b.GetBookmarkByID(ctx, user, id)
}

func (b BookmarkUseCase) DeleteBookmark(ctx context.Context, user *model.User, id int64) error {
	return b.bookmarkRepo.DeleteBookmark(ctx, user, id)
}
