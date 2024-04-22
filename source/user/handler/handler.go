package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/viniokamoto/go-store/internal/environment/server/api"
	"github.com/viniokamoto/go-store/source/user/domain/models"
	"github.com/viniokamoto/go-store/source/user/services"
)

type UserHandlerInterface interface {
	GetUsers(c *gin.Context)
	GetProfile(c *gin.Context)
	UpdateProfile(c *gin.Context)
	DeleteProfile(c *gin.Context)
}

type UserHandler struct {
	service services.UserServices
}

func UserHandlerFactory(service services.UserServices) UserHandlerInterface {
	return &UserHandler{service: service}
}

func (handler *UserHandler) GetUsers(c *gin.Context) {

	users, exception := handler.service.GetUsers()

	if exception != nil {
		api.HandleError(c, *exception)
		return
	}

	api.OkResponse(c, users, "Users found successfully")
}

func (handler *UserHandler) GetProfile(c *gin.Context) {
	id := api.UserId(c)

	user, exception := handler.service.GetUserByID(id)

	if exception != nil {
		api.HandleError(c, *exception)
		return
	}

	api.OkResponse(c, user, "User found successfully")

}

func (handler *UserHandler) UpdateProfile(c *gin.Context) {
	id := api.UserId(c)
	var request models.UserUpdateRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		api.AbortBadRequest(c, err)
		return
	}

	user, exception := handler.service.UpdateUser(id, request)

	if exception != nil {
		api.HandleError(c, *exception)
		return
	}

	api.OkResponse(c, user, "User updated successfully")

}

func (handler *UserHandler) DeleteProfile(c *gin.Context) {
	id := api.UserId(c)
	exception := handler.service.DeleteUser(id)

	if exception != nil {
		api.HandleError(c, *exception)
		return
	}

	api.OkResponse(c, nil, "User deleted successfully")
}
