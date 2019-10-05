package bookmark

import "errors"

var (
	ErrUserNotFound = errors.New("user not found")
	ErrTodoNotFound = errors.New("bookmark not found")
)