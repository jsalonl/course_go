package handler

import (
	"encoding/json"
	"net/http"
	"sample-api/model"
)

type UserHandler struct {
	userService UserService
}

func NewUserHandler(userService UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (u *UserHandler) AddUser(w http.ResponseWriter, r *http.Request) {
	// idPath := r.URL.Path // users/{id}
	// idQuery := r.URL.Query().Get("id") // users?id=12312
	// Parse request JSON to struct
	var userRequest *model.UserRequest
	err := json.NewDecoder(r.Body).Decode(&userRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// Call to service
	user, err := u.userService.Add(userRequest)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_ = json.NewEncoder(w).Encode("error")
	}

	// Write Ok
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(user)
}

func (u *UserHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
	users, err := u.userService.List()
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_ = json.NewEncoder(w).Encode("error")
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(users)
}

// UserService define signature to func
type UserService interface {
	Add(*model.UserRequest) (*model.User, error)
	List() ([]*model.User, error)
}
