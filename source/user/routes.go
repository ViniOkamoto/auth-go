package user

import (
	"github.com/gin-gonic/gin"
	"github.com/viniokamoto/go-store/internal/environment/server"
	"github.com/viniokamoto/go-store/source/user/handler"
)

func BindRoutes() []server.ApiRoute {
	return []server.ApiRoute{
		GetUsers(),
		GetProfile(),
		UpdateProfile(),
		DeleteProfile(),
	}
}

func GetUsers() server.ApiRoute {
	return server.ApiRoute{
		Path:       "/user",
		MethodType: server.GET,
		IsAdmin:    true,
		Handler: func(c *gin.Context) {
			handler := handler.CreateUserHandler()
			handler.GetUsers(c)
		},
	}
}

func GetProfile() server.ApiRoute {
	return server.ApiRoute{
		Path:       "/profile",
		MethodType: server.GET,
		Handler: func(c *gin.Context) {
			handler := handler.CreateUserHandler()
			handler.GetProfile(c)
		},
	}
}

func UpdateProfile() server.ApiRoute {
	return server.ApiRoute{
		Path:       "/profile",
		MethodType: server.PUT,
		Handler: func(c *gin.Context) {
			handler := handler.CreateUserHandler()
			handler.UpdateProfile(c)
		},
	}

}

func DeleteProfile() server.ApiRoute {
	return server.ApiRoute{
		Path:       "/profile",
		MethodType: server.DELETE,
		Handler: func(c *gin.Context) {
			handler := handler.CreateUserHandler()
			handler.DeleteProfile(c)
		},
	}
}
