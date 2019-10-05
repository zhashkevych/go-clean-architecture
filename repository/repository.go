package repository

import "context"
import "go-clean-architecture/models"

type Repository interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUser(ctx context.Context, id int64) (*models.User, error)
	CreateTodo(ctx context.Context, todo *models.Todo) error
	GetTodoByUserID(ctx context.Context, userID int64) (*models.Todo, error)
	DeleteTodo(ctx context.Context, id int64) error
}