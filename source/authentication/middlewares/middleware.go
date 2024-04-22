package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/viniokamoto/go-store/internal/environment/logging"
	"github.com/viniokamoto/go-store/internal/environment/server/api"
	"github.com/viniokamoto/go-store/source/authentication/jwt"
)

type AccessResult struct {
	AccessToken      string `json:"access_token"`
	ExpiresIn        int    `json:"expires_in"`
	RefreshExpiresIn int    `json:"refresh_expires_in"`
	RefreshToken     string `json:"refresh_token"`
	TokenType        string `json:"token_type"`
	NotBeforePolicy  int    `json:"not-before-policy"`
	SessionState     string `json:"session_state"`
	Scope            string `json:"scope"`
}

func AuthMiddleware(ctx *gin.Context) {
	reqToken, err := api.BearerToken(ctx)

	if err != nil {
		api.AbortUnauthorized(ctx)
		return
	}

	jwt := jwt.Instance

	claims, err := jwt.ValidateToken(reqToken)

	if err != nil {
		logging.Info("Error parsing token: " + err.Error())

		api.RemoveClaims(ctx)
		api.AbortUnauthorized(ctx)
		return
	}

	api.SetClaims(ctx, claims)

	ctx.Next()
}

func CustomerMiddleware(ctx *gin.Context) {
	claims := api.GetClaims(ctx)

	if claims.IsCustomer() {
		ctx.Next()
		return
	}

	api.AbortForbidden(ctx)
}

func AdminMiddleware(ctx *gin.Context) {
	claims := api.GetClaims(ctx)

	if claims.IsAdmin() {
		ctx.Next()
		return
	}

	api.AbortForbidden(ctx)
}

func StoreMiddleware(ctx *gin.Context) {
	claims := api.GetClaims(ctx)

	if claims.IsStore() {
		ctx.Next()
		return
	}

	api.AbortForbidden(ctx)
}
