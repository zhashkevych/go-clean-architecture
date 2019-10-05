package usecase

import (
	"context"
	"github.com/stretchr/testify/mock"
	"github.com/zhashkevych/go-clean-architecture/bookmark/model"
)

type BookmarkUseCaseMock struct {
	mock.Mock
}

func (m BookmarkUseCaseMock) CreateUser(ctx context.Context, user *model.User) error {
	args := m.Called(user)

	return args.Error(0)
}

func (m BookmarkUseCaseMock) GetUser(ctx context.Context, id int64) (*model.User, error) {
	args := m.Called(id)

	return args.Get(0).(*model.User), args.Error(0)
}

func (m BookmarkUseCaseMock) CreateBookmark(ctx context.Context, user *model.User, todo *model.Bookmark) error {
	args := m.Called(user, todo)

	return args.Error(0)
}

func (m BookmarkUseCaseMock) GetBookmarks(ctx context.Context, user *model.User) ([]*model.Bookmark, error) {
	args := m.Called(user)

	return args.Get(0).([]*model.Bookmark), args.Error(1)
}

func (m BookmarkUseCaseMock) GetBookmarkByID(ctx context.Context, user *model.User, id int64) (*model.Bookmark, error) {
	args := m.Called(user, id)

	return args.Get(0).(*model.Bookmark), args.Error(1)
}

func (m BookmarkUseCaseMock) DeleteBookmark(ctx context.Context, user *model.User, id int64) error {
	args := m.Called(user, id)

	return args.Error(0)
}
