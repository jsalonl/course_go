package main

import (
	"encoding/json"
	"github.com/google/uuid"
	"log"
	"net/http"
	"os"
)

//TIP To run your code, right-click the code and select <b>Run</b>. Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.

func main() {
	// Create Handler
	mux := http.NewServeMux()
	// Assign handlers or controllers
	mux.HandleFunc("POST /users", addUser)
	mux.HandleFunc("GET /users", listUsers)

	port := os.Getenv("PORT")

	// Config the server
	server := &http.Server{
		Addr:    port,
		Handler: mux,
	}

	// Start the server
	log.Fatal(server.ListenAndServe())
}

func addUser(w http.ResponseWriter, r *http.Request) {
	// idPath := r.URL.Path // users/{id}
	// idQuery := r.URL.Query().Get("id") // users?id=12312
	// Parse request JSON to struct
	var userRequest UserRequest
	err := json.NewDecoder(r.Body).Decode(&userRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Create model user
	user := User{
		ID:   uuid.New().String(),
		Name: userRequest.Name,
	}
	// Add user to users list
	usersList = append(usersList, user)

	// Write response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(user)
}

func listUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(usersList)
}

// UserRequest request in handler
type UserRequest struct {
	Name string `json:"name"`
}

// User model
type User struct {
	ID   string
	Name string
}

// List of users, to simulate repository
var usersList []User = make([]User, 0)
