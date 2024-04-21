package authentication

import (
	"github.com/gin-gonic/gin"
	"github.com/viniokamoto/go-store/internal/environment/server"
)

func BindRoutes() []server.ApiRoute {
	return []server.ApiRoute{
		Authenticate(),
		Logout(),
	}
}

func Authenticate() server.ApiRoute {
	return server.ApiRoute{
		Path:        "/authenticate",
		MethodType:  server.POST,
		IsAnonymous: true,
		Handler: func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Hello from authentication route",
			})
		},
	}

}

func Logout() server.ApiRoute {
	return server.ApiRoute{
		Path:        "/logout",
		MethodType:  server.POST,
		IsAnonymous: true,
		Handler: func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Hello from logout route",
			})
		},
	}
}
