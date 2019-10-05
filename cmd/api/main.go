package main

import (
	"github.com/zhashkevych/go-clean-architecture/server"
	"log"
)

func main() {
	app := server.NewApp("", "8000")

	if err := app.Run(); err != nil {
		log.Fatalf("%s", err.Error())
	}
}