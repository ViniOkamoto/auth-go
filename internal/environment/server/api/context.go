package api

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
)

func BearerToken(ctx *gin.Context) (string, error) {
	reqToken := ctx.GetHeader("Authorization")
	tokenSlice := strings.Split(reqToken, " ")

	if len(tokenSlice) < 2 {
		return "", errors.New("invalid authorization header format")
	} else {
		return tokenSlice[1], nil
	}
}
