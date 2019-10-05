package repository

import (
	"context"
	"github.com/stretchr/testify/mock"
	"github.com/zhashkevych/go-clean-architecture/todo/model"
)

type StorageMock struct {
	mock.Mock
}

func (s *StorageMock) CreateUser(ctx context.Context, user *model.User) error {
	args := s.Called(user)

	return args.Error(0)
}

func (s *StorageMock) GetUser(ctx context.Context, id int64) (*model.User, error) {
	args := s.Called(id)

	return args.Get(0).(*model.User), args.Error(1)
}

func (s *StorageMock) CreateTodo(ctx context.Context, todo *model.Todo) error {
	args := s.Called(todo)

	return args.Error(0)
}

func (s *StorageMock) GetTodosByUserID(ctx context.Context, userID int64) ([]*model.Todo, error) {
	args := s.Called(userID)

	return args.Get(0).([]*model.Todo), args.Error(1)
}

func (s *StorageMock) GetTodoByID(ctx context.Context, id int64) (*model.Todo, error) {
	args := s.Called(id)

	return args.Get(0).(*model.Todo), args.Error(1)
}

func (s *StorageMock) DeleteTodo(ctx context.Context, id int64) error {
	args := s.Called(id)

	return args.Error(0)
}
