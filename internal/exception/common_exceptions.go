package internal

import (
	"net/http"
)

func DBException() *Exception {
	return &Exception{
		Code:    http.StatusInternalServerError,
		Message: "Error connecting to database",
	}
}

func ServerException() *Exception {
	return &Exception{
		Code:    http.StatusInternalServerError,
		Message: "Server error",
	}
}
