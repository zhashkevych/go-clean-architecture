package localstorage

import (
	"context"
	"github.com/zhashkevych/go-clean-architecture/auth"
	"go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
	"sync"
)

type UserLocalStorage struct {
	users map[uuid.UUID]*auth.User
	mutex *sync.Mutex
}

func NewUserLocalStorage() *UserLocalStorage {
	return &UserLocalStorage{
		users: make(map[uuid.UUID]*auth.User),
		mutex: new(sync.Mutex),
	}
}

func (s *UserLocalStorage) CreateUser(ctx context.Context, user *auth.User) error {
	s.mutex.Lock()
	s.users[user.ID] = user
	s.mutex.Unlock()

	return nil
}

func (s *UserLocalStorage) GetUser(ctx context.Context, username, password string) (*auth.User, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	for _, user := range s.users {
		if user.Username == username && user.Password == password {
			return user, nil
		}
	}

	return nil, auth.ErrUserNotFound
}
