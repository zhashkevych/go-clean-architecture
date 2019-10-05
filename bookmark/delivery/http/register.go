package http

import (
	"github.com/gin-gonic/gin"
	"github.com/zhashkevych/go-clean-architecture/bookmark"
	"github.com/zhashkevych/go-clean-architecture/bookmark/delivery/http/handler"
)

func RegisterHTTPEndpoints(router *gin.RouterGroup, uc bookmark.UseCase) {
	h := handler.NewHandler(uc)

	noauth := router.Group("/")
	{
		noauth.POST("/sign-up", h.SignUp)
		noauth.POST("/sign-in", h.SignIn)
	}

	bookmarks := router.Group("/bookmarks")
	{
		bookmarks.POST("", h.Create)
		bookmarks.GET("", h.Get)
		bookmarks.GET(":id", h.Get)
		bookmarks.DELETE(":id", h.Delete)
	}
}
