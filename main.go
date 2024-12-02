package main

import (
	"log"
	"net/http"
	"os"
	"sample-api/handler"
	"sample-api/repository"
	"sample-api/routes"
	"sample-api/service"
)

//TIP To run your code, right-click the code and select <b>Run</b>. Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.

func main() {
	// Create Handler
	mux := http.NewServeMux()
	// Dependencies injection
	userRepository := repository.NewUserRepository()
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
