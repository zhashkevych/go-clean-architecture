package handler

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

func (h *Handler) SignUp(c *gin.Context) {

}

func (h *Handler) SignIn(c *gin.Context) {

}