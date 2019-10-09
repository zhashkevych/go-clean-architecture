package auth

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `bson:"_id"`
	Username string    `bson:"username"`
	Password string    `bson:"password"`
}
