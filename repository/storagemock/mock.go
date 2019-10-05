package storagemock

import (
	"context"
	"github.com/stretchr/testify/mock"
	"github.com/zhashkevych/go-clean-architecture/models"
)

type StorageMock struct {
	mock.Mock
}

func (s *StorageMock) CreateUser(ctx context.Context, user *models.User) error {
	args := s.Called(user)

	return args.Error(0)
}

func (s *StorageMock) GetUser(ctx context.Context, id int64) (*models.User, error) {
	args := s.Called(id)

	return args.Get(0).(*models.User), args.Error(1)
}

func (s *StorageMock) CreateTodo(ctx context.Context, todo *models.Todo) error {
	args := s.Called(todo)

	return args.Error(0)
}

func (s *StorageMock) GetTodosByUserID(ctx context.Context, userID int64) ([]*models.Todo, error) {
	args := s.Called(userID)

	return args.Get(0).([]*models.Todo), args.Error(1)
}

func (s *StorageMock) GetTodoByID(ctx context.Context, id int64) (*models.Todo, error) {
	args := s.Called(id)

	return args.Get(0).(*models.Todo), args.Error(1)
}

func (s *StorageMock) DeleteTodo(ctx context.Context, id int64) error {
	args := s.Called(id)

	return args.Error(0)
}
