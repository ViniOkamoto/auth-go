package api

import (
	"github.com/gin-gonic/gin"
	"github.com/viniokamoto/go-store/source/authentication/jwt"
)

func RemoveClaims(ctx *gin.Context) {
	ctx.Set("claims", nil)
}

func SetClaims(ctx *gin.Context, claims jwt.TokenClaims) {
	ctx.Set("claims", claims)
}

func GetClaims(ctx *gin.Context) jwt.TokenClaims {
	claims, exists := ctx.Get("claims")

	if exists {
		return claims.(jwt.TokenClaims)
	} else {
		return jwt.TokenClaims{}
	}
}

func UserId(ctx *gin.Context) string {
	return GetClaims(ctx).Subject
}
