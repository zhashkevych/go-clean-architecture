package auth

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID
	Username string
	Password string
}
