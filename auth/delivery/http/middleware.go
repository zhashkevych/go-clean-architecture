package http

import (
	"github.com/gin-gonic/gin"
	"github.com/zhashkevych/go-clean-architecture/auth"
	"net/http"
	"strings"
)

type AuthMiddleware struct {
	usecase auth.UseCase
}

func NewAuthMiddleware(usecase auth.UseCase) gin.HandlerFunc {
	return (&AuthMiddleware{
		usecase: usecase,
	}).Handle
}

func (m *AuthMiddleware) Handle(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	if headerParts[0] != "Bearer" {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	user, err := m.usecase.ParseToken(c.Request.Context(), headerParts[1])
	if err != nil {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	c.Set(auth.CtxUserKey, user)
}
