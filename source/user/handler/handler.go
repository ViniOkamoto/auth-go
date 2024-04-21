package handler

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/viniokamoto/go-store/internal/environment/server/api"
	"github.com/viniokamoto/go-store/internal/utils/validator"
	"github.com/viniokamoto/go-store/source/user/domain/models"
	"github.com/viniokamoto/go-store/source/user/services"
)

type UserHandlerInterface interface {
	GetUsers(c *gin.Context)
	GetUserByID(c *gin.Context)
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
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

	api.OkResponse(c, users)
}

func (handler *UserHandler) GetUserByID(c *gin.Context) {
	id := c.Param("id")

	idUint, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		api.AbortBadRequest(c, errors.New("invalid id"), nil)
		return
	}

	user, exception := handler.service.GetUserByID(uint(idUint))

	if exception != nil {
		api.HandleError(c, *exception)
		return
	}

	api.OkResponse(c, user)

}

func (handler *UserHandler) CreateUser(c *gin.Context) {
	var request models.UserCreateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		api.AbortBadRequest(c, err, nil)
		return
	}

	validationErrors := validator.ValidateStruct(request)
	if validationErrors != nil {
		api.AbortBadRequest(c, errors.New("invalid request"), validationErrors)
		return
	}

	exception := handler.service.CreateUser(request)
	if exception != nil {
		api.HandleError(c, *exception)
		return
	}

	api.OkResponse(c, nil)

}

func (handler *UserHandler) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var request models.UserUpdateRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		api.AbortBadRequest(c, err)
		return
	}

	idUint, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		api.AbortBadRequest(c, errors.New("invalid id"))
		return
	}

	user, exception := handler.service.UpdateUser(uint(idUint), request)

	if exception != nil {
		api.HandleError(c, *exception)
		return
	}

	api.OkResponse(c, user)

}

func (handler *UserHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")

	idUint, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		api.AbortBadRequest(c, errors.New("invalid id"))
		return
	}

	exception := handler.service.DeleteUser(uint(idUint))

	if exception != nil {
		api.HandleError(c, *exception)
		return
	}

	api.OkResponse(c, nil)
}
