package routes

import (
	"net/http"
	"sample-api/handler"
)

func RegisterUserRoutes(mux *http.ServeMux, handler *handler.UserHandler) {
	mux.HandleFunc("POST /users", handler.AddUser)
	mux.HandleFunc("GET /users", handler.ListUsers)
}
