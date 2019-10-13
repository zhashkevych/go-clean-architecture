package bookmark

import "go.mongodb.org/mongo-driver/x/mongo/driver/uuid"

type Bookmark struct {
	ID     uuid.UUID `bson:"_id"`
	UserID uuid.UUID `bson:"userId"`
	URL    string    `bson:"url"`
	Title  string    `bson:"title"`
}
