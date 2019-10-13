package localstorage

import (
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/zhashkevych/go-clean-architecture/auth"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

func TestGetUser(t *testing.T) {
	s := NewUserLocalStorage()

	id1 := primitive.NewObjectID()

	user := &auth.User{
		ID:       id1,
		Username: "user",
		Password: "password",
	}

	err := s.CreateUser(context.Background(), user)
	assert.NoError(t, err)

	returnedUser, err := s.GetUser(context.Background(), "user", "password")
	assert.NoError(t, err)
	assert.Equal(t, user, returnedUser)

	returnedUser, err = s.GetUser(context.Background(), "user", "")
	assert.Error(t, err)
	assert.Equal(t, err, auth.ErrUserNotFound)
}
