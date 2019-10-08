package localcache

import (
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/zhashkevych/go-clean-architecture/auth"
	"github.com/zhashkevych/go-clean-architecture/bookmark"
	"testing"
)

func TestGetBookmarks(t *testing.T) {
	id := uuid.New()
	user := &auth.User{ID: id}

	s := NewBookmarkLocalStorage()

	for i := 0; i < 10; i++ {
		bm := &bookmark.Bookmark{
			ID:     id,
			UserID: user.ID,
		}

		err := s.CreateBookmark(context.Background(), user, bm)
		assert.NoError(t, err)
	}

	returnedBookmarks, err := s.GetBookmarks(context.Background(), user)
	assert.NoError(t, err)

	assert.Equal(t, 10, len(returnedBookmarks))
}

func TestDeleteBookmark(t *testing.T) {
	id1 := uuid.New()
	id2 := uuid.New()

	user1 := &auth.User{ID: id1}
	user2 := &auth.User{ID: id2}

	bmID := uuid.New()
	bm := &bookmark.Bookmark{ID: bmID, UserID: user1.ID}

	s := NewBookmarkLocalStorage()

	err := s.CreateBookmark(context.Background(), user1, bm)
	assert.NoError(t, err)

	err = s.DeleteBookmark(context.Background(), user1, bmID)
	assert.NoError(t, err)

	err = s.CreateBookmark(context.Background(), user1, bm)
	assert.NoError(t, err)

	err = s.DeleteBookmark(context.Background(), user2, bmID)
	assert.Error(t, err)
	assert.Equal(t, err, bookmark.ErrBookmarkNotFound)
}
