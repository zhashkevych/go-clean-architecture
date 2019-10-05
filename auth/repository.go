package auth

import (
	"context"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *User) error
	GetUser(ctx context.Context, username, password string) (*User, error)
}
