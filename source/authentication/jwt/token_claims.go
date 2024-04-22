package jwt

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	"github.com/viniokamoto/go-store/internal/environment/logging"
)

type TokenClaims struct {
	RoleID uint `json:"roleId"`
	jwt.RegisteredClaims
}

func (t TokenClaims) IsAdmin() bool {
	logging.Info(fmt.Sprint("RoleID: ", t.RoleID))
	return t.RoleID == 1
}

func (t TokenClaims) IsStore() bool {
	return t.RoleID == 2
}

func (t TokenClaims) IsCustomer() bool {
	return t.RoleID == 3
}
