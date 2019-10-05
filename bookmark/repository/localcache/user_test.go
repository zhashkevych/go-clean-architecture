package localcache

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/zhashkevych/go-clean-architecture/bookmark"
	"github.com/zhashkevych/go-clean-architecture/bookmark/model"
	"testing"
)

func TestGetUser(t *testing.T) {
	s := NewUserLocalStorage()

	user := &model.User{
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
