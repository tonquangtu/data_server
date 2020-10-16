package service

import (
	"data_server/internal/model"
	"data_server/internal/repository"
)

type UserService interface {
	GetUser(id int) (*model.User, error)
	AddUser(user *model.User)
}

type userService struct {
	userRepository *repository.UserRepository
}

func NewUserService(userRepository *repository.UserRepository) UserService {
	return &userService{userRepository: userRepository}
}

func (us *userService) GetUser(id int) (*model.User, error) {
	return us.userRepository.GetUser(id)
}

func (us *userService) AddUser(user *model.User) {
	us.userRepository.AddUser(user)
}