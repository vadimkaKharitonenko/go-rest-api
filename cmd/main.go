package main

import (
	"log"

	todo "github.com/vadimkaKharitonenko/go-rest-api"
	"github.com/vadimkaKharitonenko/go-rest-api/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(todo.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error accured while running http server: %s", err.Error())
	}
}
