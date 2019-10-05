package localcache

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/zhashkevych/go-clean-architecture/auth"
	"github.com/zhashkevych/go-clean-architecture/bookmark"
	"testing"
)

func TestGetBookmarks(t *testing.T) {
	user := &auth.User{ID: 1}

	s := NewBookmarkLocalStorage()

	for i := 0; i < 10; i++ {
		bm := &bookmark.Bookmark{
			ID:     int64(i),
			UserID: user.ID,
		}

		err := s.CreateBookmark(context.Background(), user, bm)
		assert.NoError(t, err)
	}

	returnedBookmarks, err := s.GetBookmarks(context.Background(), user)
	assert.NoError(t, err)

	assert.Equal(t, 10, len(returnedBookmarks))
}

func TestGetBookmarkByID(t *testing.T) {
	user := &auth.User{ID: 1}
	user2 := &auth.User{ID: 2}
	bm := &bookmark.Bookmark{ID: 15, UserID: user.ID}

	s := NewBookmarkLocalStorage()

	err := s.CreateBookmark(context.Background(), user, bm)
	assert.NoError(t, err)

	returnedTodo, err := s.GetBookmarkByID(context.Background(), user, int64(15))
	assert.NoError(t, err)
	assert.Equal(t, bm, returnedTodo)

	returnedTodo, err = s.GetBookmarkByID(context.Background(), user, int64(0))
	assert.Error(t, err)
	assert.Equal(t, err, bookmark.ErrBookmarkNotFound)

	returnedTodo, err = s.GetBookmarkByID(context.Background(), user2, int64(15))
	assert.Error(t, err)
	assert.Equal(t, err, bookmark.ErrBookmarkNotFound)
}

func TestDeleteBookmark(t *testing.T) {
	user := &auth.User{ID: 1}
	user2 := &auth.User{ID: 2}
	bm := &bookmark.Bookmark{ID: 15, UserID: user.ID}

	s := NewBookmarkLocalStorage()

	err := s.CreateBookmark(context.Background(), user, bm)
	assert.NoError(t, err)

	err = s.DeleteBookmark(context.Background(), user, int64(15))
	assert.NoError(t, err)

	_, err = s.GetBookmarkByID(context.Background(), user, bm.ID)
	assert.Error(t, err)
	assert.Equal(t, err, bookmark.ErrBookmarkNotFound)

	err = s.CreateBookmark(context.Background(), user, bm)
	assert.NoError(t, err)

	err = s.DeleteBookmark(context.Background(), user2, int64(15))
	assert.Error(t, err)
	assert.Equal(t, err, bookmark.ErrBookmarkNotFound)
}
