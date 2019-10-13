package mock

import (
	"context"
	"github.com/stretchr/testify/mock"
	"github.com/zhashkevych/go-clean-architecture/auth"
	"github.com/zhashkevych/go-clean-architecture/bookmark"
)

type BookmarkStorageMock struct {
	mock.Mock
}

func (s *BookmarkStorageMock) CreateBookmark(ctx context.Context, user *auth.User, bm *bookmark.Bookmark) error {
	args := s.Called(user, bm)

	return args.Error(0)
}

func (s *BookmarkStorageMock) GetBookmarks(ctx context.Context, user *auth.User) ([]*bookmark.Bookmark, error) {
	args := s.Called(user)

	return args.Get(0).([]*bookmark.Bookmark), args.Error(1)
}

func (s *BookmarkStorageMock) DeleteBookmark(ctx context.Context, user *auth.User, id string) error {
	args := s.Called(user, id)

	return args.Error(0)
}
