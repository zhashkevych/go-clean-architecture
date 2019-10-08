package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"

	"github.com/zhashkevych/go-clean-architecture/auth"
	"github.com/zhashkevych/go-clean-architecture/bookmark"

	authhttp "github.com/zhashkevych/go-clean-architecture/auth/delivery/http"
	authlocalcache "github.com/zhashkevych/go-clean-architecture/auth/repository/localstorage"
	authusecase "github.com/zhashkevych/go-clean-architecture/auth/usecase"
	bmhttp "github.com/zhashkevych/go-clean-architecture/bookmark/delivery/http"
	bmlocalcache "github.com/zhashkevych/go-clean-architecture/bookmark/repository/localcache"
	bmusecase "github.com/zhashkevych/go-clean-architecture/bookmark/usecase"
)

type App struct {
	httpServer *http.Server

	bookmarkUC bookmark.UseCase
	authUC     auth.UseCase
}

func NewApp(hashSalt, signingKey string) *App {
	userRepo := authlocalcache.NewUserLocalStorage()
	bookmarkRepo := bmlocalcache.NewBookmarkLocalStorage()

	return &App{
		bookmarkUC: bmusecase.NewBookmarkUseCase(bookmarkRepo),
		authUC:     authusecase.NewAuthUseCase(userRepo, hashSalt, []byte(signingKey)),
	}
}

func (a *App) Run(port string) error {
	// Init gin handler
	router := gin.Default()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)

	// Set up http handlers
	// SignUp/SignIn endpoints
	authhttp.RegisterHTTPEndpoints(router, a.authUC)

	// API endpoints
	authMiddleware := authhttp.NewAuthMiddleware(a.authUC)
	api := router.Group("/api", authMiddleware)

	bmhttp.RegisterHTTPEndpoints(api, a.bookmarkUC)

	// HTTP Server
	a.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if err := a.httpServer.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
