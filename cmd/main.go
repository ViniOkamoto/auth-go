package main

import (
	"github.com/viniokamoto/go-store/internal/environment"
	"github.com/viniokamoto/go-store/internal/environment/database"
	"github.com/viniokamoto/go-store/internal/environment/logging"
	"github.com/viniokamoto/go-store/internal/environment/server"
	"github.com/viniokamoto/go-store/internal/environment/server/routes"
	"github.com/viniokamoto/go-store/internal/utils/validator"
	"github.com/viniokamoto/go-store/source/authentication/jwt"
)

func main() {

	validator.Init()

	err := environment.Init()
	if err != nil {
		logging.FatalObject(err)
	}

	jwt.Init()

	database.InitDBConnection()

	server := server.CreateServer(server.Options{})
	server.AddRoutes(routes.BindServerRoutes())

	server.Start(environment.Config.Port)

}
