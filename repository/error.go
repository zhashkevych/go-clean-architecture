package repository

import "errors"

var (
	ErrUserNotFound = errors.New("user not found")
	ErrTodoNotFound = errors.New("todo not found")
)