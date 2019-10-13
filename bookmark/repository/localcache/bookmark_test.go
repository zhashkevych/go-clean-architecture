package localcache

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/zhashkevych/go-clean-architecture/bookmark"
	"github.com/zhashkevych/go-clean-architecture/models"
	"testing"
)

func TestGetBookmarks(t *testing.T) {
	id := "id"
	user := &models.User{ID: id}

	s := NewBookmarkLocalStorage()

	for i := 0; i < 10; i++ {
		bm := &models.Bookmark{
			ID:     fmt.Sprintf("id%d", i),
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
	id1 := "id1"
	id2 := "id2"

	user1 := &models.User{ID: id1}
	user2 := &models.User{ID: id2}

	bmID := "bmID"
	bm := &models.Bookmark{ID: bmID, UserID: user1.ID}

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
