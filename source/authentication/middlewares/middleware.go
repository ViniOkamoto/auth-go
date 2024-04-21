package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/viniokamoto/go-store/internal/environment/logging"
	"github.com/viniokamoto/go-store/internal/environment/server/api"
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

	parser := jwt.Parser{}

	// Note:
	// This does not verify the token using the signing signature.
	// This is safe as long as this request is forwarded from a gateway which handles the actual verificiation
	token, _, err := parser.ParseUnverified(reqToken, jwt.MapClaims{})

	if err != nil {
		logging.Info("Error parsing token: " + err.Error())
		api.AbortUnauthorized(ctx)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		api.SetClaims(ctx, claims)
	}
}
