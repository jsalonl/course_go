package service

import (
	"sample-api/handler"
	"sample-api/model"
)

type userService struct {
	userRepository UserRepository
}

func NewUserService(repository UserRepository) handler.UserService {
	return &userService{
		repository,
	}
}

func (u *userService) Add(request *model.UserRequest) (*model.User, error) {
	user := model.NewUser(request.Name)

	return u.userRepository.Add(user)
}

func (u *userService) List() ([]*model.User, error) {
	return u.userRepository.List()
}

// UserRepository is the interface to func
type UserRepository interface {
	Add(user *model.User) (*model.User, error)
	List() ([]*model.User, error)
}
