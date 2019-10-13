package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/zhashkevych/go-clean-architecture/auth"
	"github.com/zhashkevych/go-clean-architecture/bookmark"

	authhttp "github.com/zhashkevych/go-clean-architecture/auth/delivery/http"
	authmongo "github.com/zhashkevych/go-clean-architecture/auth/repository/mongo"
	authusecase "github.com/zhashkevych/go-clean-architecture/auth/usecase"
	bmhttp "github.com/zhashkevych/go-clean-architecture/bookmark/delivery/http"
	bmmongo "github.com/zhashkevych/go-clean-architecture/bookmark/repository/mongo"
	bmusecase "github.com/zhashkevych/go-clean-architecture/bookmark/usecase"
)

type App struct {
	httpServer *http.Server

	bookmarkUC bookmark.UseCase
	authUC     auth.UseCase
}

func NewApp() *App {
	db := initDB()

	userRepo := authmongo.NewUserRepository(db, viper.GetString("mongo.user_collection"))
	bookmarkRepo := bmmongo.NewBookmarkRepository(db, viper.GetString("mongo.bookmark_collection"))

	return &App{
		bookmarkUC: bmusecase.NewBookmarkUseCase(bookmarkRepo),
		authUC: authusecase.NewAuthUseCase(
			userRepo,
			viper.GetString("auth.hash_salt"),
			[]byte(viper.GetString("auth.signing_key")),
			viper.GetDuration("auth.token_ttl"),
		),
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

	go func() {
		if err := a.httpServer.ListenAndServe(); err != nil {
			log.Fatalf("Failed to listen and serve: %+v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return a.httpServer.Shutdown(ctx)
}

func initDB() *mongo.Database {
	client, err := mongo.NewClient(options.Client().ApplyURI(viper.GetString("mongo.uri")))
	if err != nil {
		log.Fatalf("Error occured while establishing connection to mongoDB")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	return client.Database(viper.GetString("mongo.name"))
}
