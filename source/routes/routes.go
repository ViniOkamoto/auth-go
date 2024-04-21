package routes

import (
	"github.com/viniokamoto/go-store/internal/environment/server"
	"github.com/viniokamoto/go-store/source/authentication"
	"github.com/viniokamoto/go-store/source/user"
)

func BindServerRoutes() []server.ApiRoute {
	authenticationRoutes := authentication.BindRoutes()
	userRoutes := user.BindRoutes()

	return append(authenticationRoutes, userRoutes...)
}
