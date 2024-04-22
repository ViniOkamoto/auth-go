package exceptions

import (
	"net/http"

	internal "github.com/viniokamoto/go-store/internal/exception"
)

func UserNotFoundException() *internal.Exception {
	return &internal.Exception{
		Code:    http.StatusNotFound,
		Message: "User does not exist or password is invalid",
	}
}

func InvalidPasswordException() *internal.Exception {
	return &internal.Exception{
		Code:    http.StatusBadRequest,
		Message: "User does not exist or password is invalid",
	}
}

func InvalidTokenException() *internal.Exception {
	return &internal.Exception{
		Code:    http.StatusUnauthorized,
		Message: "Invalid token",
	}
}
