package di

import (
	"gin-starter/internal/application/service"
	"gin-starter/internal/infrastructure/config"
	"gin-starter/internal/infrastructure/repository_impl"
	"gin-starter/internal/interface/handler"
)

type Locator struct {
	UserHandler *handler.UserHandler
}

func InitLocator() *Locator {
	database := config.GetDatabase()

	userRepository := repository_impl.NewUserRepositoryImpl(database)

	userService := service.NewUserService(userRepository)

	userHandler := handler.NewUserHandler(userService)

	return &Locator{
		UserHandler: userHandler,
	}
}
