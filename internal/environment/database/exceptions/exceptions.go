package exceptions

import (
	"net/http"

	internal "github.com/viniokamoto/go-store/internal/exception"
)

func DBException() *internal.Exception {
	return &internal.Exception{
		Code:    http.StatusInternalServerError,
		Message: "Error connecting to database",
	}
}
