package mock

import (
	"context"
	"github.com/stretchr/testify/mock"
	"github.com/zhashkevych/go-clean-architecture/models"
)

type BookmarkStorageMock struct {
	mock.Mock
}

func (s *BookmarkStorageMock) CreateBookmark(ctx context.Context, user *models.User, bm *models.Bookmark) error {
	args := s.Called(user, bm)

	return args.Error(0)
}

func (s *BookmarkStorageMock) GetBookmarks(ctx context.Context, user *models.User) ([]*models.Bookmark, error) {
	args := s.Called(user)

	return args.Get(0).([]*models.Bookmark), args.Error(1)
}

func (s *BookmarkStorageMock) DeleteBookmark(ctx context.Context, user *models.User, id string) error {
	args := s.Called(user, id)

	return args.Error(0)
}
