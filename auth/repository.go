package auth

import (
	"context"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *User) error
	GetUser(ctx context.Context, id int64) (*User, error)
}
