package user

import (
	"github.com/gin-gonic/gin"
	"github.com/viniokamoto/go-store/internal/environment/server"
	"github.com/viniokamoto/go-store/source/user/handler"
)

func BindRoutes() []server.ApiRoute {
	return []server.ApiRoute{
		GetUser(),
		GetUserByID(),
		CreateUser(),
		UpdateUser(),
		DeleteUser(),
	}
}

func GetUser() server.ApiRoute {
	return server.ApiRoute{
		Path:        "/user",
		MethodType:  server.GET,
		IsAnonymous: true,
		Handler: func(c *gin.Context) {
			handler := handler.CreateUserHandler()
			handler.GetUsers(c)
		},
	}
}

func GetUserByID() server.ApiRoute {
	return server.ApiRoute{
		Path:        "/user/:id",
		MethodType:  server.GET,
		IsAnonymous: true,
		Handler: func(c *gin.Context) {
			handler := handler.CreateUserHandler()
			handler.GetUserByID(c)
		},
	}
}

func CreateUser() server.ApiRoute {
	return server.ApiRoute{
		Path:        "/user",
		MethodType:  server.POST,
		IsAnonymous: true,
		Handler: func(c *gin.Context) {
			handler := handler.CreateUserHandler()
			handler.CreateUser(c)
		},
	}
}

func UpdateUser() server.ApiRoute {
	return server.ApiRoute{
		Path:        "/user/:id",
		MethodType:  server.PUT,
		IsAnonymous: true,
		Handler: func(c *gin.Context) {
			handler := handler.CreateUserHandler()
			handler.UpdateUser(c)
		},
	}

}

func DeleteUser() server.ApiRoute {
	return server.ApiRoute{
		Path:        "/user/:id",
		MethodType:  server.DELETE,
		IsAnonymous: true,
		Handler: func(c *gin.Context) {
			handler := handler.CreateUserHandler()
			handler.DeleteUser(c)
		},
	}
}
