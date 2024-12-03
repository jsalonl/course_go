package service

import (
	"fmt"
	"sample-api/internal/errors_app"
	"sample-api/internal/handler"
	"sample-api/internal/model"
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

	existByName, err := u.userRepository.ExistByName(user.Name)
	if err != nil {
		return nil, fmt.Errorf("error in db %v", err.Error())
	}
	if existByName {
		return nil, errors_app.UserAlreadyExist
	}

	return u.userRepository.Add(user)
}

func (u *userService) List() ([]*model.User, error) {
	return u.userRepository.List()
}

// UserRepository is the interface to func
type UserRepository interface {
	Add(user *model.User) (*model.User, error)
	List() ([]*model.User, error)
	ExistByName(name string) (bool, error)
}
