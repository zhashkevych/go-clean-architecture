package todo

import (
	"context"
	"github.com/zhashkevych/go-clean-architecture/todo/model"
)

type Repository interface {
	CreateUser(ctx context.Context, user *model.User) error
	GetUser(ctx context.Context, id int64) (*model.User, error)
	CreateTodo(ctx context.Context, todo *model.Todo) error
	GetTodosByUserID(ctx context.Context, userID int64) ([]*model.Todo, error)
	GetTodoByID(ctx context.Context, id int64) (*model.Todo, error)
	DeleteTodo(ctx context.Context, id int64) error
}
