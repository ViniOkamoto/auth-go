package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/viniokamoto/go-store/internal/environment/logging"
	internal "github.com/viniokamoto/go-store/internal/exception"
	"github.com/viniokamoto/go-store/internal/utils"
)

type ApiError struct {
	Message string          `json:"message"`
	Path    string          `json:"path"`
	Code    int             `json:"code"`
	Detail  json.RawMessage `json:"detail"`
}

type ApiResponse struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Error   *ApiError   `json:"error"`
}

func (r ApiResponse) ToString() string {
	return fmt.Sprintf("Data: %T\nError: %T", r.Data, r.Error)
}

func (r ApiResponse) Serialize() (*string, *string) {
	return utils.ToJson(r)
}

func NewError(message, path string, statusCode int) *ApiError {
	result := ApiError{}
	result.Message = message
	result.Path = path
	result.Code = statusCode
	return &result
}

func HandleError(c *gin.Context, exception internal.Exception) {

	resp := ApiResponse{}

	resp.Error = NewError(exception.Message, c.Request.RequestURI, exception.Code)

	c.JSON(http.StatusBadRequest, resp)
	c.Abort()
}

func AbortBadRequest(c *gin.Context, err error, details ...json.RawMessage) {

	var detail json.RawMessage
	if len(details) > 0 {
		detail = details[0]
	}

	resp := ApiResponse{
		Error: &ApiError{
			Message: err.Error(),
			Code:    400,
			Path:    c.Request.RequestURI,
			Detail:  detail,
		},
	}
	c.JSON(http.StatusBadRequest, resp)
	c.Abort()
}

func AbortUnauthorized(c *gin.Context) {
	resp := ApiResponse{
		Error: &ApiError{
			Message: "Unauthorized",
			Code:    401,
			Path:    c.Request.RequestURI,
		},
	}
	c.JSON(http.StatusUnauthorized, resp)
	c.Abort()
}

func AbortForbidden(c *gin.Context) {
	err := ApiError{
		Message: "Forbidden",
		Code:    403,
		Path:    c.Request.RequestURI,
	}
	resp := ApiResponse{Error: &err}

	c.JSON(http.StatusForbidden, resp)
	c.Abort()
}

func AbortInternalServerError(c *gin.Context, err error) {
	logging.Error("Internal Server Error: %s", err)

	apiError := ApiError{
		Message: "Internal Server Error",
		Code:    500,
		Path:    c.Request.RequestURI,
	}
	resp := ApiResponse{
		Error: &apiError,
	}

	c.JSON(http.StatusInternalServerError, resp)
	c.Abort()
}

func AbortNotFound(c *gin.Context) {
	resp := ApiError{
		Message: "Not Found",
		Code:    404,
		Path:    c.Request.RequestURI,
	}

	c.JSON(http.StatusNotFound, resp)
	c.Abort()
}

func OkResponse(c *gin.Context, data interface{}, message ...string) {
	resp := ApiResponse{
		Data: data,
	}

	if len(message) > 0 {
		resp.Message = message[0]
	}

	c.JSON(http.StatusOK, resp)
}

func NoContentResponse(c *gin.Context) {
	c.Status(http.StatusNoContent)
}
