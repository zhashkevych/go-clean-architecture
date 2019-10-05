package bookmark

import "errors"

var (
	ErrUserNotFound     = errors.New("user not found")
	ErrBookmarkNotFound = errors.New("bookmark not found")
)
