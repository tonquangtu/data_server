package internal

import (
	"data_server/internal/config"
	"data_server/internal/repository"
	"data_server/internal/service"
)

type Container struct {
	Configuration config.Configuration
	UserService service.UserService
}

func NewContainer() (*Container, error) {
	configuration := config.Config()
	db, err := config.CreateDB(configuration.DBConfiguration)
	if err != nil {
		return nil, err
	}

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)

	return &Container{
		Configuration: configuration,
		UserService: userService,
	}, nil
}