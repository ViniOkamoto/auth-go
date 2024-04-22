package handler

import (
	"github.com/viniokamoto/go-store/internal/environment/database"
	"github.com/viniokamoto/go-store/source/authentication/services"
	userRepository "github.com/viniokamoto/go-store/source/user/repository"

	userService "github.com/viniokamoto/go-store/source/user/services"
)

func CreateAuthenticationHandler() AuthHandlerInterface {
	userRepository := userRepository.UserRepositoryFactory(database.DB)

	userService := userService.UserServicesFactory(userRepository)

	service := services.AuthServiceFactory(userRepository, userService)

	return AuthHandlerFactory(service)

}
