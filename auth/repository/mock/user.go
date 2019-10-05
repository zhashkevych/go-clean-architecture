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

func (s *UserStorageMock) GetUser(ctx context.Context, username, password string) (*auth.User, error) {
	args := s.Called(username, password)

	return args.Get(0).(*auth.User), args.Error(1)
}