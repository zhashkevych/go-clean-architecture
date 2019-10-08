package bookmark

import "github.com/google/uuid"

type Bookmark struct {
	ID          uuid.UUID
	UserID      uuid.UUID
	URL 		string
	Title       string
}
