package main

import (
	"log"
	"net/http"
	"os"
	"sample-api/internal/config"
	"sample-api/internal/handler"
	"sample-api/internal/infrastructure/storage/postgres"
	"sample-api/internal/repository"
	"sample-api/internal/routes"
	"sample-api/internal/service"
)

//TIP To run your code, right-click the code and select <b>Run</b>. Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.

func main() {
	// Init dependencies
	db := postgres.Connect(config.DatabaseConfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "develop",
		Password: "develop",
		Database: "develop",
	})
	// Create Handler
	mux := http.NewServeMux()
	// Dependencies injection
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	routes.RegisterUserRoutes(mux, userHandler)

	port := os.Getenv("PORT")

	// Config the server
	server := &http.Server{
		Addr:    port,
		Handler: mux,
	}

	// Start the server
	log.Fatal(server.ListenAndServe())
}