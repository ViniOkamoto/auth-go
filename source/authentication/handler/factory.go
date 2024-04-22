package handler

import (
	"github.com/viniokamoto/go-store/internal/environment/database"
	"github.com/viniokamoto/go-store/source/authentication/services"
	"github.com/viniokamoto/go-store/source/user/repository"

	userService "github.com/viniokamoto/go-store/source/user/services"
)

func CreateAuthenticationHandler() AuthHandlerInterface {
	repository := repository.UserRepositoryFactory(database.DB)

	userService := userService.UserServicesFactory(repository)

	service := services.AuthServiceFactory(repository, userService)

	return AuthHandlerFactory(service)

}
