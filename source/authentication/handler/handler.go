package handler

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/viniokamoto/go-store/internal/environment/server/api"
	"github.com/viniokamoto/go-store/internal/utils/validator"
	"github.com/viniokamoto/go-store/source/authentication/domain/models"
	"github.com/viniokamoto/go-store/source/authentication/services"
)

type AuthHandlerInterface interface {
	Login(c *gin.Context)
	Register(c *gin.Context)
}

type AuthHandler struct {
	service services.AuthServiceInterface
}

func AuthHandlerFactory(authService services.AuthServiceInterface) AuthHandlerInterface {
	return &AuthHandler{service: authService}

}

func (handler *AuthHandler) Login(c *gin.Context) {
	var request models.AuthenticationRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		api.AbortBadRequest(c, errors.New("invalid request"), nil)
		return
	}

	if exception := validator.ValidateStruct(request); exception != nil {
		api.AbortBadRequest(c, errors.New("invalid request"), exception)
		return
	}

	response, exception := handler.service.Login(request)

	if exception != nil {
		api.HandleError(c, *exception)
		return
	}

	api.OkResponse(c, response, "User logged in successfully")
}

func (handler *AuthHandler) Register(c *gin.Context) {
	var request models.RegisterRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		api.AbortBadRequest(c, errors.New("invalid request"), nil)
		return
	}

	if exception := validator.ValidateStruct(request); exception != nil {
		api.AbortBadRequest(c, errors.New("invalid request"), exception)
		return
	}

	exception := handler.service.Register(request)

	if exception != nil {
		api.HandleError(c, *exception)
		return
	}

	api.OkResponse(c, nil, "User created successfully")
}
