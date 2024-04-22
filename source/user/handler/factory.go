package handler

import (
	"github.com/viniokamoto/go-store/internal/environment/database"
	"github.com/viniokamoto/go-store/source/user/repository"
	"github.com/viniokamoto/go-store/source/user/services"
)

func CreateUserHandler() UserHandlerInterface {

	repository := repository.UserRepositoryFactory(database.DB)

	service := services.UserServicesFactory(repository)

	return UserHandlerFactory(service)

}
