package server

import (
	"github.com/gin-gonic/gin"

	"github.com/zhashkevych/go-clean-architecture/auth"
	"github.com/zhashkevych/go-clean-architecture/bookmark"

	authhttp "github.com/zhashkevych/go-clean-architecture/auth/delivery/http"
	authlocalcache "github.com/zhashkevych/go-clean-architecture/auth/repository/localstorage"
	authusecase "github.com/zhashkevych/go-clean-architecture/auth/usecase"
	bmhttp "github.com/zhashkevych/go-clean-architecture/bookmark/delivery/http"
	bmlocalcache "github.com/zhashkevych/go-clean-architecture/bookmark/repository/localcache"
	bmusecase "github.com/zhashkevych/go-clean-architecture/bookmark/usecase"
)

// TODO; use http.Client
type App struct {
	httpAddr string
	httpPort string

	bookmarkUC bookmark.UseCase
	authUC     auth.UseCase
}

// TODO: inject hashSalt and secret
func NewApp(addr, port string) *App {
	userRepo := authlocalcache.NewUserLocalStorage()
	bookmarkRepo := bmlocalcache.NewBookmarkLocalStorage()

	return &App{
		bookmarkUC: bmusecase.NewBookmarkUseCase(bookmarkRepo),
		authUC:     authusecase.NewAuthUseCase(userRepo, "hashSalt", []byte("secret")),
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

	authhttp.RegisterHTTPEndpoints(router, a.authUC)

	authMiddleware := authhttp.NewAuthMiddleware(a.authUC)
	api := router.Group("/api", authMiddleware)

	bmhttp.RegisterHTTPEndpoints(api, a.bookmarkUC)

	if err := router.Run(a.httpAddr + ":" + a.httpPort); err != nil {
		return err
	}

	return nil
}
