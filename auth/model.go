package auth

import "go.mongodb.org/mongo-driver/x/mongo/driver/uuid"

type User struct {
	ID       uuid.UUID `bson:"_id"`
	Username string    `bson:"username"`
	Password string    `bson:"password"`
}
