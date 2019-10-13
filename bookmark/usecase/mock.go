package usecase

import (
	"context"
	"github.com/stretchr/testify/mock"
	"github.com/zhashkevych/go-clean-architecture/auth"
	"github.com/zhashkevych/go-clean-architecture/bookmark"
)

type BookmarkUseCaseMock struct {
	mock.Mock
}

func (m BookmarkUseCaseMock) CreateBookmark(ctx context.Context, user *auth.User, url, title string) error {
	args := m.Called(user, url, title)

	return args.Error(0)
}

func (m BookmarkUseCaseMock) GetBookmarks(ctx context.Context, user *auth.User) ([]*bookmark.Bookmark, error) {
	args := m.Called(user)

	return args.Get(0).([]*bookmark.Bookmark), args.Error(1)
}

func (m BookmarkUseCaseMock) DeleteBookmark(ctx context.Context, user *auth.User, id string) error {
	args := m.Called(user, id)

	return args.Error(0)
}
