package mongo

import (
	"context"
	"github.com/zhashkevych/go-clean-architecture/auth"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	db *mongo.Collection
}

func NewUserRepository(db *mongo.Database, collection string) *UserRepository {
	return &UserRepository{
		db: db.Collection(collection),
	}
}

func (r UserRepository) CreateUser(ctx context.Context, user *auth.User) error {
	res, err := r.db.InsertOne(ctx, user)
	if err != nil {
		return err
	}

	user.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

func (r UserRepository) GetUser(ctx context.Context, username, password string) (*auth.User, error) {
	user := new(auth.User)
	err := r.db.FindOne(ctx, bson.M{
		"username": username,
		"password": password,
	}).Decode(user)

	if err != nil {
		return nil, err
	}

	return user, nil
}
