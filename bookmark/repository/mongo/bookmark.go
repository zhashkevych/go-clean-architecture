package mongo

import (
	"context"
	"github.com/zhashkevych/go-clean-architecture/auth"
	"github.com/zhashkevych/go-clean-architecture/bookmark"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
)

type BookmarkRepository struct {
	db *mongo.Collection
}

func NewBookmarkRepository(db *mongo.Database, collection string) *BookmarkRepository {
	return &BookmarkRepository{
		db: db.Collection(collection),
	}
}

func (r BookmarkRepository) CreateBookmark(ctx context.Context, user *auth.User, bm *bookmark.Bookmark) error {
	bm.UserID = user.ID

	res, err := r.db.InsertOne(ctx, bm)
	if err != nil {
		return err
	}

	bm.ID = res.InsertedID.(uuid.UUID)
	return nil
}

func (r BookmarkRepository) GetBookmarks(ctx context.Context, user *auth.User) ([]*bookmark.Bookmark, error) {
	cur, err := r.db.Find(ctx, bson.M{
		"userId": user.ID,
	})
	defer cur.Close(ctx)

	if err != nil {
		return nil, err
	}

	out := make([]*bookmark.Bookmark, 0)

	for cur.Next(ctx) {
		user := new(bookmark.Bookmark)
		err := cur.Decode(user)
		if err != nil {
			return nil, err
		}

		out = append(out, user)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}

	return out, nil
}

func (r BookmarkRepository) DeleteBookmark(ctx context.Context, user *auth.User, id uuid.UUID) error {
	_, err := r.db.DeleteOne(ctx, bson.M{"_id": id, "userId": user.ID})
	return err
}
