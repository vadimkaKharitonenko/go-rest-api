package main

import (
	"log"

	"github.com/spf13/viper"
	todo "github.com/vadimkaKharitonenko/go-rest-api"
	"github.com/vadimkaKharitonenko/go-rest-api/pkg/handler"
	"github.com/vadimkaKharitonenko/go-rest-api/pkg/repository"
	"github.com/vadimkaKharitonenko/go-rest-api/pkg/service"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing config: %s", err.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error accured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
