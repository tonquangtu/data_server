package handler

import (
	"github.com/tonquangtu/data_server/internal/service"
	"github.com/tonquangtu/data_server/pkg/rpc/dto"
)

type UserHandler struct {
	userService service.UserService
}

func NewHandler(userService service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (handler *UserHandler) GetUser(id int, result *dto.GetUserDto) error {
	user, err := handler.userService.GetUser(id)
	if err != nil {
		return err
	}
	result = user
	return nil
}

func (handler *UserHandler) AddUser(userDto dto.AddUserDto, result *dto.AddUserDto) error {
	err := handler.userService.AddUser(&userDto)
	if err != nil {
		return nil
	}
	result = &userDto
	return nil
}
