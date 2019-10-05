package localcache

import (
	"context"
	"sync"

	"github.com/zhashkevych/go-clean-architecture/models"
	"github.com/zhashkevych/go-clean-architecture/repository"
)

type LocalStorage struct {
	users map[int64]*models.User
	um    *sync.Mutex

	todos map[int64]*models.Todo
	tm    *sync.Mutex
}

func NewLocalStorage() *LocalStorage {
	return &LocalStorage{
		users: make(map[int64]*models.User),
		todos: make(map[int64]*models.Todo),
		um:    new(sync.Mutex),
		tm:    new(sync.Mutex),
	}
}

func (s *LocalStorage) CreateUser(ctx context.Context, user *models.User) error {
	s.um.Lock()
	s.users[user.ID] = user
	s.um.Unlock()

	return nil
}

func (s *LocalStorage) GetUser(ctx context.Context, id int64) (*models.User, error) {
	s.um.Lock()
	defer s.um.Unlock()
	if user, ex := s.users[id]; ex {
		return user, nil
	}

	return nil, repository.ErrUserNotFound
}

func (s *LocalStorage) CreateTodo(ctx context.Context, todo *models.Todo) error {
	s.tm.Lock()
	s.todos[todo.ID] = todo
	s.tm.Unlock()

	return nil
}

func (s *LocalStorage) GetTodosByUserID(ctx context.Context, userID int64) ([]*models.Todo, error) {
	todos := make([]*models.Todo, 0)

	s.tm.Lock()
	for _, todo := range s.todos {
		if todo.UserID == userID {
			todos = append(todos, todo)
		}
	}
	s.tm.Unlock()

	return todos, nil
}

func (s *LocalStorage) GetTodoByID(ctx context.Context, id int64) (*models.Todo, error) {
	s.tm.Lock()
	defer s.tm.Unlock()
	if todo, ex := s.todos[id]; ex {
		return todo, nil
	}

	return nil, repository.ErrTodoNotFound
}

func (s *LocalStorage) DeleteTodo(ctx context.Context, id int64) error {
	s.tm.Lock()
	delete(s.todos, id)
	s.tm.Unlock()

	return nil
}
