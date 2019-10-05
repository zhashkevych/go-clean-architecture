package localcache

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/zhashkevych/go-clean-architecture/bookmark/model"
	"testing"
)

func TestGetUser(t *testing.T) {
	s := NewLocalStorage()

	user := &model.User{
		ID:       0,
		Username: "user",
	}

	err := s.CreateUser(context.Background(), user)
	assert.NoError(t, err)

	returnedUser, err := s.GetUser(context.Background(), 0)
	assert.NoError(t, err)

	assert.Equal(t, user, returnedUser)
}

func TestGetTodosByUserID(t *testing.T) {
	userID := int64(10)
	userTodos := make([]*model.Bookmark, 0)

	s := NewLocalStorage()

	for i := 0; i <= 10; i++ {
		todo := &model.Bookmark{
			ID: int64(i),
			UserID: func(i int) int64 {
				if i%2 == 0 {
					return userID
				}

				return int64(0)
			}(i),
		}

		if i%2 == 0 {
			userTodos = append(userTodos, todo)
		}

		err := s.CreateTodo(context.Background(), todo)
		assert.NoError(t, err)
	}

	returnedTodos, err := s.GetTodosByUserID(context.Background(), userID)
	assert.NoError(t, err)

	assert.Equal(t, len(userTodos), len(returnedTodos))
}

func TestGetTodoByID(t *testing.T) {
	todo := &model.Bookmark{ID: 15}

	s := NewLocalStorage()

	err := s.CreateTodo(context.Background(), todo)
	assert.NoError(t, err)

	returnedTodo, err := s.GetTodoByID(context.Background(), int64(15))
	assert.NoError(t, err)
	assert.Equal(t, todo, returnedTodo)

	returnedTodo, err = s.GetTodoByID(context.Background(), int64(0))
	assert.Error(t, err)
}
