package localcache

import (
	"context"
	"github.com/zhashkevych/go-clean-architecture/bookmark"
	"sync"

	"github.com/zhashkevych/go-clean-architecture/bookmark/model"
)

type LocalStorage struct {
	users map[int64]*model.User
	um    *sync.Mutex

	todos map[int64]*model.Bookmark
	tm    *sync.Mutex
}

func NewLocalStorage() *LocalStorage {
	return &LocalStorage{
		users: make(map[int64]*model.User),
		todos: make(map[int64]*model.Bookmark),
		um:    new(sync.Mutex),
		tm:    new(sync.Mutex),
	}
}

func (s *LocalStorage) CreateUser(ctx context.Context, user *model.User) error {
	s.um.Lock()
	s.users[user.ID] = user
	s.um.Unlock()

	return nil
}

func (s *LocalStorage) GetUser(ctx context.Context, id int64) (*model.User, error) {
	s.um.Lock()
	defer s.um.Unlock()
	if user, ex := s.users[id]; ex {
		return user, nil
	}

	return nil, bookmark.ErrUserNotFound
}

func (s *LocalStorage) CreateTodo(ctx context.Context, todo *model.Bookmark) error {
	s.tm.Lock()
	s.todos[todo.ID] = todo
	s.tm.Unlock()

	return nil
}

func (s *LocalStorage) UpdateTodo(ctx context.Context, id int64, todo *model.Bookmark) error {
	s.tm.Lock()
	s.todos[id] = todo
	s.tm.Unlock()

	return nil
}

func (s *LocalStorage) GetTodosByUserID(ctx context.Context, userID int64) ([]*model.Bookmark, error) {
	todos := make([]*model.Bookmark, 0)

	s.tm.Lock()
	for _, todo := range s.todos {
		if todo.UserID == userID {
			todos = append(todos, todo)
		}
	}
	s.tm.Unlock()

	return todos, nil
}

func (s *LocalStorage) GetTodoByID(ctx context.Context, id int64) (*model.Bookmark, error) {
	s.tm.Lock()
	defer s.tm.Unlock()
	if todo, ex := s.todos[id]; ex {
		return todo, nil
	}

	return nil, bookmark.ErrTodoNotFound
}

func (s *LocalStorage) DeleteTodo(ctx context.Context, id int64) error {
	s.tm.Lock()
	delete(s.todos, id)
	s.tm.Unlock()

	return nil
}
