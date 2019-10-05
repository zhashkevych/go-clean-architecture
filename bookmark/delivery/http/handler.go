package http

import (
	"github.com/gin-gonic/gin"
	"github.com/zhashkevych/go-clean-architecture/bookmark"
)

type Handler struct {
	useCase bookmark.UseCase
}

func NewHandler(useCase bookmark.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

func (h *Handler) Create(c *gin.Context) {

}

func (h *Handler) Get(c *gin.Context) {

}

func (h *Handler) Delete(c *gin.Context) {

}
