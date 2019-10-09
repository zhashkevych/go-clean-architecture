package bookmark

import "github.com/google/uuid"

type Bookmark struct {
	ID     uuid.UUID `bson:"_id"`
	UserID uuid.UUID `bson:"userId"`
	URL    string    `bson:"url"`
	Title  string    `bson:"title"`
}
