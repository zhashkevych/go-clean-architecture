package http

import (
	"github.com/gin-gonic/gin"
	"github.com/zhashkevych/go-clean-architecture/bookmark"
)

func RegisterHTTPEndpoints(router *gin.RouterGroup, uc bookmark.UseCase) {
	h := NewHandler(uc)

	bookmarks := router.Group("/bookmarks")
	{
		bookmarks.POST("", h.Create)
		bookmarks.GET("", h.Get)
		bookmarks.DELETE("", h.Delete)
	}
}
