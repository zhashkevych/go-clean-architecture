package mock

import (
	"context"
	"github.com/stretchr/testify/mock"
	"github.com/zhashkevych/go-clean-architecture/bookmark/model"
)

type UserStorageMock struct {
	mock.Mock
}

func (s *UserStorageMock) CreateUser(ctx context.Context, user *model.User) error {
	args := s.Called(user)

	return args.Error(0)
}

func (s *UserStorageMock) GetUser(ctx context.Context, id int64) (*model.User, error) {
	args := s.Called(id)

	return args.Get(0).(*model.User), args.Error(1)
}
