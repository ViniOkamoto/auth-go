package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/viniokamoto/go-store/internal/environment"
	"github.com/viniokamoto/go-store/internal/environment/logging"
)

var Instance *JWT

type JWT struct {
	key []byte
}

type JWTInterface interface {
	GenerateAccessToken(userId string, roleId uint) (string, int, error)
	ValidateToken(token string) (TokenClaims, error)
}

func Init() {
	config := environment.Config
	key := config.JWTKey

	Instance = &JWT{[]byte(key)}
}

func (j *JWT) GenerateAccessToken(userId string, roleId uint) (string, int, error) {
	logging.Info("Generating access token for user: " + userId)
	logging.Info(fmt.Sprint("RoleID: ", roleId))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, TokenClaims{
		RoleID: roleId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 4)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Subject:   userId,
		},
	})

	tokenString, err := token.SignedString(j.key)

	if err != nil {
		return "", 0, err
	}

	expiresIn := int(time.Now().Add(time.Hour * 4).Unix())

	return tokenString, expiresIn, nil
}

func (j *JWT) parseToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return j.key, nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil

}

func (j *JWT) ValidateToken(tokenString string) (TokenClaims, error) {

	token, err := j.parseToken(tokenString)

	if claims, ok := token.Claims.(TokenClaims); ok && token.Valid {
		return claims, nil
	} else {
		return TokenClaims{}, err
	}
}
