package mock

import (
	"context"
	"github.com/stretchr/testify/mock"
	"github.com/zhashkevych/go-clean-architecture/bookmark/model"
)

type BookmarkStorageMock struct {
	mock.Mock
}

func (s *BookmarkStorageMock) CreateBookmark(ctx context.Context, todo *model.Bookmark) error {
	args := s.Called(todo)

	return args.Error(0)
}

func (s *BookmarkStorageMock) UpdateBookmark(ctx context.Context, id int64, todo *model.Bookmark) error {
	args := s.Called(todo)

	return args.Error(0)
}

func (s *BookmarkStorageMock) GetBookmarksByUserID(ctx context.Context, userID int64) ([]*model.Bookmark, error) {
	args := s.Called(userID)

	return args.Get(0).([]*model.Bookmark), args.Error(1)
}

func (s *BookmarkStorageMock) GetBookmarkByID(ctx context.Context, id int64) (*model.Bookmark, error) {
	args := s.Called(id)

	return args.Get(0).(*model.Bookmark), args.Error(1)
}

func (s *BookmarkStorageMock) DeleteBookmark(ctx context.Context, id int64) error {
	args := s.Called(id)

	return args.Error(0)
}
