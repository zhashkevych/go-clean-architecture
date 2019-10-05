package server

import (
	"github.com/gin-gonic/gin"
	"github.com/zhashkevych/go-clean-architecture/bookmark"
	deliveryhttp "github.com/zhashkevych/go-clean-architecture/bookmark/delivery/http"
	"github.com/zhashkevych/go-clean-architecture/bookmark/repository/localcache"
	"github.com/zhashkevych/go-clean-architecture/bookmark/usecase"
)

type App struct {
	httpAddr   string
	httpPort   string
	bookmarkUC bookmark.UseCase
}

func NewApp(addr, port string) *App {
	userRepo := localcache.NewUserLocalStorage()
	bookmarkRepo := localcache.NewBookmarkLocalStorage()

	return &App{
		bookmarkUC: usecase.NewBookmarkUseCase(userRepo, bookmarkRepo),
		httpAddr:   addr,
		httpPort:   port,
	}
}

func (a *App) Run() error {
	router := gin.Default()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)

	api := router.Group("/api")

	deliveryhttp.RegisterHTTPEndpoints(api, a.bookmarkUC)

	if err := router.Run(a.httpAddr + ":" + a.httpPort); err != nil {
		return err
	}

	return nil
}