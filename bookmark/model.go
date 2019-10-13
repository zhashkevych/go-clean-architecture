package bookmark

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Bookmark struct {
	ID     primitive.ObjectID `bson:"_id"`
	UserID primitive.ObjectID `bson:"userId"`
	URL    string             `bson:"url"`
	Title  string             `bson:"title"`
}