package middlewares

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/viniokamoto/go-store/source/domain/exceptions"
)

type ErrorResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Error   any    `json:"error"`
}

func ExceptionMiddleware(c *gin.Context, recovered interface{}) {
	if except, ok := recovered.(*exceptions.HttpException); ok {
		errorResponse := ErrorResponse{
			Code:    except.StatusCode,
			Message: except.Message,
			Error:   "Something went wrong during processing, try again later",
		}
		c.JSON(except.StatusCode, errorResponse)
	} else {
		log.Printf("Exception not mapped: %s", recovered)
		errorResponse := ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Internal server error",
			Error:   recovered,
		}
		c.JSON(http.StatusInternalServerError, errorResponse)
	}
}
