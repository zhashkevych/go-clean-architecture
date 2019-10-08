package main

import (
	"github.com/zhashkevych/go-clean-architecture/config"
	"github.com/zhashkevych/go-clean-architecture/server"
	"log"
)

func main() {
	cfg, err := config.Init()
	if err != nil {
		log.Fatalf("%s", err.Error())
	}

	app := server.NewApp(cfg.Auth.HashSalt, cfg.Auth.SigningKey)

	if err := app.Run(cfg.Port); err != nil {
		log.Fatalf("%s", err.Error())
	}
}
