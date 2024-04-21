package api

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func SetClaims(ctx *gin.Context, claims jwt.MapClaims) {
	ctx.Set("claims", claims)
}
