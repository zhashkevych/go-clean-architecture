package mongo

import (
	"context"
	"github.com/zhashkevych/go-clean-architecture/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Bookmark struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	UserID primitive.ObjectID `bson:"userId"`
	URL    string             `bson:"url"`
	Title  string             `bson:"title"`
}

type BookmarkRepository struct {
	db *mongo.Collection
}

func NewBookmarkRepository(db *mongo.Database, collection string) *BookmarkRepository {
	return &BookmarkRepository{
		db: db.Collection(collection),
	}
}

func (r BookmarkRepository) CreateBookmark(ctx context.Context, user *models.User, bm *models.Bookmark) error {
	bm.UserID = user.ID

	model := toModel(bm)

	res, err := r.db.InsertOne(ctx, model)
	if err != nil {
		return err
	}

	bm.ID = res.InsertedID.(primitive.ObjectID).Hex()
	return nil
}

func (r BookmarkRepository) GetBookmarks(ctx context.Context, user *models.User) ([]*models.Bookmark, error) {
	uid, _ := primitive.ObjectIDFromHex(user.ID)
	cur, err := r.db.Find(ctx, bson.M{
		"userId": uid,
	})
	defer cur.Close(ctx)

	if err != nil {
		return nil, err
	}

	out := make([]*Bookmark, 0)

	for cur.Next(ctx) {
		user := new(Bookmark)
		err := cur.Decode(user)
		if err != nil {
			return nil, err
		}

		out = append(out, user)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}

	return toBookmarks(out), nil
}

func (r BookmarkRepository) DeleteBookmark(ctx context.Context, user *models.User, id string) error {
	objID, _ := primitive.ObjectIDFromHex(id)
	uID, _ := primitive.ObjectIDFromHex(user.ID)

	_, err := r.db.DeleteOne(ctx, bson.M{"_id": objID, "userId": uID})
	return err
}

func toModel(b *models.Bookmark) *Bookmark {
	uid, _ := primitive.ObjectIDFromHex(b.UserID)

	return &Bookmark{
		UserID: uid,
		URL:    b.URL,
		Title:  b.Title,
	}
}

func toBookmark(b *Bookmark) *models.Bookmark {
	return &models.Bookmark{
		ID:     b.ID.Hex(),
		UserID: b.UserID.Hex(),
		URL:    b.URL,
		Title:  b.Title,
	}
}

func toBookmarks(bs []*Bookmark) []*models.Bookmark {
	out := make([]*models.Bookmark, len(bs))

	for i, b := range bs {
		out[i] = toBookmark(b)
	}

	return out
}
