package mock

import (
	"context"
	"github.com/stretchr/testify/mock"
	"github.com/zhashkevych/go-clean-architecture/auth"
)

type UserStorageMock struct {
	mock.Mock
}

func (s *UserStorageMock) CreateUser(ctx context.Context, user *auth.User) error {
	args := s.Called(user)

	return args.Error(0)
}

func (s *UserStorageMock) GetUser(ctx context.Context, id int64) (*auth.User, error) {
	args := s.Called(id)

	return args.Get(0).(*auth.User), args.Error(1)
}
