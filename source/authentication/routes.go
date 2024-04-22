package authentication

import (
	"github.com/gin-gonic/gin"
	"github.com/viniokamoto/go-store/internal/environment/server"
	"github.com/viniokamoto/go-store/source/authentication/handler"
)

func BindRoutes() []server.ApiRoute {
	return []server.ApiRoute{
		Login(),
		Register(),
	}
}

func Login() server.ApiRoute {
	return server.ApiRoute{
		Path:        "/login",
		MethodType:  server.POST,
		IsAnonymous: true,
		Handler: func(c *gin.Context) {
			handler := handler.CreateAuthenticationHandler()

			handler.Login(c)
		},
	}

}

func Register() server.ApiRoute {
	return server.ApiRoute{
		Path:        "/register",
		MethodType:  server.POST,
		IsAnonymous: true,
		Handler: func(c *gin.Context) {
			handler := handler.CreateAuthenticationHandler()

			handler.Register(c)
		},
	}

}
