package http

import (
	"github.com/gin-gonic/gin"
	"github.com/zhashkevych/go-clean-architecture/auth"
)

type Handler struct {
	useCase auth.UseCase
}

func NewHandler(useCase auth.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

func (h *Handler) SignUp(c *gin.Context) {

}

func (h *Handler) SignIn(c *gin.Context) {

}
