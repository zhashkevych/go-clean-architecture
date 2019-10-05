package todo

import "errors"

var (
	ErrUserNotFound = errors.New("user not found")
	ErrTodoNotFound = errors.New("todo not found")
)