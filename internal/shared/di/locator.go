package di

import (
	"gin-starter/internal/application/service"
	"gin-starter/internal/domain/repository"
	"gin-starter/internal/infrastructure/config"
	"gin-starter/internal/infrastructure/repository_impl"
	"gin-starter/internal/interface/handler"
)

type Locator struct {
	UserRepository repository.UserRepository

	UserService *service.UserService

	UserHandler *handler.UserHandler
}

func InitLocator() *Locator {
	database := config.GetDatabase()
	redisClient := config.GetRedisClient()

	userRepository := repository_impl.NewUserRepositoryImpl(database)

	userService := service.NewUserService(userRepository)

	userHandler := handler.NewUserHandler(redisClient, userService)

	return &Locator{
		UserRepository: userRepository,
		UserService:    userService,
		UserHandler:    userHandler,
	}
}
