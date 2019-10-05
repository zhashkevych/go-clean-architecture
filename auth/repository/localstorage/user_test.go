package localstorage

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/zhashkevych/go-clean-architecture/auth"
	"github.com/zhashkevych/go-clean-architecture/bookmark"
	"testing"
)

func TestGetUser(t *testing.T) {
	s := NewUserLocalStorage()

	user := &auth.User{
		ID:       0,
		Username: "user",
	}

	err := s.CreateUser(context.Background(), user)
	assert.NoError(t, err)

	returnedUser, err := s.GetUser(context.Background(), 0)
	assert.NoError(t, err)
	assert.Equal(t, user, returnedUser)

	returnedUser, err = s.GetUser(context.Background(), 1)
	assert.Error(t, err)
	assert.Equal(t, err, bookmark.ErrUserNotFound)
}
