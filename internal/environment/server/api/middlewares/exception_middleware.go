package middlewares

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/viniokamoto/go-store/internal/environment/server/api"
	internal "github.com/viniokamoto/go-store/internal/exception"
)

func ExceptionMiddleware(c *gin.Context, recovered interface{}) {
	if except, ok := recovered.(*internal.Exception); ok {
		errorResponse := api.ApiError{
			Code:    except.Code,
			Message: except.Message,
			Path:    c.Request.RequestURI,
		}
		c.JSON(except.Code, errorResponse)
	} else {
		log.Printf("Exception not mapped: %s", recovered)
		errorResponse := api.ApiError{
			Code:    http.StatusInternalServerError,
			Message: "Internal server error",
			Path:    c.Request.RequestURI,
		}
		c.JSON(http.StatusInternalServerError, errorResponse)
	}
}
