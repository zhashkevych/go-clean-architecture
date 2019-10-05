package localcache

import (
	"context"
	"github.com/zhashkevych/go-clean-architecture/bookmark"
	"sync"

	"github.com/zhashkevych/go-clean-architecture/bookmark/model"
)

type UserLocalStorage struct {
	users map[int64]*model.User
	mutex *sync.Mutex
}

func NewUserLocalStorage() *UserLocalStorage {
	return &UserLocalStorage{
		users: make(map[int64]*model.User),
		mutex: new(sync.Mutex),
	}
}

func (s *UserLocalStorage) CreateUser(ctx context.Context, user *model.User) error {
	s.mutex.Lock()
	s.users[user.ID] = user
	s.mutex.Unlock()

	return nil
}

func (s *UserLocalStorage) GetUser(ctx context.Context, id int64) (*model.User, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if user, ex := s.users[id]; ex {
		return user, nil
	}

	return nil, bookmark.ErrUserNotFound
}
