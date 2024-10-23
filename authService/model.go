package authService

import "github.com/golang-jwt/jwt/v4"

type MpcUserToken struct {
	UserInfo *UserInfo `json:"userInfo"`
	jwt.RegisteredClaims
}
