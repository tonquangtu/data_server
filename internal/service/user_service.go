package service

import (
	"errors"
	"github.com/tonquangtu/data_server/internal/model"
	"github.com/tonquangtu/data_server/internal/repository"
	"github.com/tonquangtu/data_server/pkg/rpc/dto"
)

type UserService interface {
	GetUser(id int) (*dto.GetUserDto, error)
	AddUser(user *dto.AddUserDto) error
}

type userService struct {
	userRepository  *repository.UserRepository
	passwordEncoder PasswordEncoder
}

func NewUserService(userRepository *repository.UserRepository, passwordEncoder PasswordEncoder) UserService {
	return &userService{
		userRepository:  userRepository,
		passwordEncoder: passwordEncoder,
	}
}

func (us *userService) GetUser(id int) (*dto.GetUserDto, error) {
	user, err := us.userRepository.GetUser(id)
	if err != nil {
		return nil, err
	}
	return &dto.GetUserDto{
		ID:           user.ID,
		Username:     user.Username,
		PasswordHash: user.PasswordHash,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
	}, nil
}

func (us *userService) AddUser(userDto *dto.AddUserDto) error {
	err := us.validateUserDto(userDto)
	if err != nil {
		return err
	}
	user := &model.User{
		Username:     userDto.Username,
		PasswordHash: us.marshalPassword(userDto.Password),
		CreatedAt:    userDto.CreatedAt,
	}
	us.userRepository.AddUser(user)
	return nil
}

func (us *userService) validateUserDto(userDto *dto.AddUserDto) error {
	if userDto == nil {
		return errors.New("user cannot be empty")
	}
	if userDto.Username == "" {
		return errors.New("username cannot be empty")
	}
	if userDto.Password == "" {
		return errors.New("password cannot be empty")
	}
	return nil
}

func (us *userService) marshalPassword(password string) string {
	return us.passwordEncoder.Encode(password)
}
