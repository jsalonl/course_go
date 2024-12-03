package model

import "github.com/google/uuid"

// UserRequest request in handler
type UserRequest struct {
	// Name represent name of user
	Name string `json:"name"`
}

// User model
type User struct {
	ID   string
	Name string
}

// NewUser constructor to user
func NewUser(name string) *User {
	return &User{
		ID:   uuid.New().String(),
		Name: name,
	}
}
