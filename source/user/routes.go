package user

import (
	"github.com/gin-gonic/gin"
	"github.com/viniokamoto/go-store/internal/environment/server"
)

func BindRoutes() []server.ApiRoute {
	return []server.ApiRoute{
		GetUser(),
	}
}

func GetUser() server.ApiRoute {
	return server.ApiRoute{
		Path:        "/user",
		MethodType:  server.GET,
		IsAnonymous: true,
		Handler: func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Hello from user route",
			})
		},
	}
}
