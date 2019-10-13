package http

import (
	"github.com/gin-gonic/gin"
	"github.com/zhashkevych/go-clean-architecture/auth"
	"github.com/zhashkevych/go-clean-architecture/bookmark"
	"github.com/zhashkevych/go-clean-architecture/models"
	"net/http"
)

type Bookmark struct {
	ID    string `json:"id"`
	URL   string `json:"url"`
	Title string `json:"title"`
}

type Handler struct {
	useCase bookmark.UseCase
}

func NewHandler(useCase bookmark.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

type createInput struct {
	URL   string `json:"url"`
	Title string `json:"title"`
}

func (h *Handler) Create(c *gin.Context) {
	inp := new(createInput)
	if err := c.BindJSON(inp); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	user := c.MustGet(auth.CtxUserKey).(*models.User)

	if err := h.useCase.CreateBookmark(c.Request.Context(), user, inp.URL, inp.Title); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

type getResponse struct {
	Bookmarks []*Bookmark `json:"bookmarks"`
}

func (h *Handler) Get(c *gin.Context) {
	user := c.MustGet(auth.CtxUserKey).(*models.User)

	bms, err := h.useCase.GetBookmarks(c.Request.Context(), user)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, &getResponse{
		Bookmarks: toBookmarks(bms),
	})
}

type deleteInput struct {
	ID string `json:"id"`
}

func (h *Handler) Delete(c *gin.Context) {
	inp := new(deleteInput)
	if err := c.BindJSON(inp); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	user := c.MustGet(auth.CtxUserKey).(*models.User)

	if err := h.useCase.DeleteBookmark(c.Request.Context(), user, inp.ID); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

func toBookmarks(bs []*models.Bookmark) []*Bookmark {
	out := make([]*Bookmark, len(bs))

	for i, b := range bs {
		out[i] = toBookmark(b)
	}

	return out
}

func toBookmark(b *models.Bookmark) *Bookmark {
	return &Bookmark{
		ID:    b.ID,
		URL:   b.URL,
		Title: b.Title,
	}
}
