package localcache

import (
	"context"
	"go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
	"sync"

	"github.com/zhashkevych/go-clean-architecture/auth"
	"github.com/zhashkevych/go-clean-architecture/bookmark"
)

type BookmarkLocalStorage struct {
	bookmarks map[uuid.UUID]*bookmark.Bookmark
	mutex     *sync.Mutex
}

func NewBookmarkLocalStorage() *BookmarkLocalStorage {
	return &BookmarkLocalStorage{
		bookmarks: make(map[uuid.UUID]*bookmark.Bookmark),
		mutex:     new(sync.Mutex),
	}
}

func (s *BookmarkLocalStorage) CreateBookmark(ctx context.Context, user *auth.User, bm *bookmark.Bookmark) error {
	bm.UserID = user.ID

	s.mutex.Lock()
	s.bookmarks[bm.ID] = bm
	s.mutex.Unlock()

	return nil
}

func (s *BookmarkLocalStorage) GetBookmarks(ctx context.Context, user *auth.User) ([]*bookmark.Bookmark, error) {
	bookmarks := make([]*bookmark.Bookmark, 0)

	s.mutex.Lock()
	for _, bm := range s.bookmarks {
		if bm.UserID == user.ID {
			bookmarks = append(bookmarks, bm)
		}
	}
	s.mutex.Unlock()

	return bookmarks, nil
}

func (s *BookmarkLocalStorage) DeleteBookmark(ctx context.Context, user *auth.User, id uuid.UUID) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	bm, ex := s.bookmarks[id]
	if ex && bm.UserID == user.ID {
		delete(s.bookmarks, id)
		return nil
	}

	return bookmark.ErrBookmarkNotFound
}
