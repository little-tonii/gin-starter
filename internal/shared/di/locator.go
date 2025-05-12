package di

import (
	"health-care-system/internal/application/service"
	"health-care-system/internal/infrastructure/config"
	"health-care-system/internal/infrastructure/repository_impl"
	"health-care-system/internal/interface/handler"
)

type Locator struct {
	UserHandler *handler.UserHandler
}

func InitLocator() *Locator {
	database := config.GetDatabase()

	userRepository := repository_impl.NewUserRepositoryImpl(database)
	roleRepository := repository_impl.NewRoleRepositoryImpl(database)

	userService := service.NewUserService(userRepository, roleRepository)

	userHandler := handler.NewUserHandler(userService)

	return &Locator{
		UserHandler: userHandler,
	}
}
