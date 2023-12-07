package main

import (
	"log"

	"github.com/3P3-21/user/internal/config"
	"github.com/3P3-21/user/internal/server"
	"github.com/3P3-21/user/internal/service"
	"github.com/3P3-21/user/internal/store/postgres"
)

func main() {
	config, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	db, err := postgres.New(config.App.PostgresDSN)
	if err != nil {
		log.Fatal(err)
	}

	userStore := postgres.NewUserStore(db)

	// Init services
	services := service.New(userStore)

	server := server.NewServer(&server.Config{
		Addr:    config.App.Addr,
		Port:    config.App.Port,
		Service: services,
	})

	server.SetupRouter()

	log.Print("Server up and running.")

	err = server.RunServer()
	if err != nil {
		log.Panic(err)
	}
}
