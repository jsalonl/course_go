package repository

import (
	"sample-api/model"
	"sample-api/service"
)

type userRepository struct {
}

func NewUserRepository() service.UserRepository {
	return &userRepository{}
}

func (u *userRepository) Add(user *model.User) (*model.User, error) {
	usersList = append(usersList, user)

	return user, nil
}

func (u *userRepository) List() ([]*model.User, error) {
	return usersList, nil
}

// usersList List of users, in memory repository
var usersList = make([]*model.User, 0)
