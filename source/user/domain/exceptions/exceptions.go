package exceptions

import (
	"net/http"

	internal "github.com/viniokamoto/go-store/internal/exception"
)

func UserNotFoundException() *internal.Exception {
	return &internal.Exception{
		Code:    http.StatusNotFound,
		Message: "User not found",
	}
}

func UserEmailAlreadyExistException() *internal.Exception {
	return &internal.Exception{
		Code:    http.StatusConflict,
		Message: "User email already exists",
	}
}

func InvalidPasswordException() *internal.Exception {
	return &internal.Exception{
		Code:    http.StatusBadRequest,
		Message: "Invalid password",
	}
}

func InvalidEmailException() *internal.Exception {
	return &internal.Exception{
		Code:    http.StatusBadRequest,
		Message: "Invalid email",
	}
}
