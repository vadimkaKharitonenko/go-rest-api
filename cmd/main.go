package main

import (
	"log"

	todo "github.com/vadimkaKharitonenko/go-rest-api"
	"github.com/vadimkaKharitonenko/go-rest-api/pkg/handler"
	"github.com/vadimkaKharitonenko/go-rest-api/pkg/repository"
	"github.com/vadimkaKharitonenko/go-rest-api/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error accured while running http server: %s", err.Error())
	}
}
