package authServiceModel

import (
	"encoding/json"

	"github.com/golang-jwt/jwt/v4"
)

type UserInfo struct {
	UserId string `json:"appPubKey"`
	AppId  string `json:"appId"`

	// TimeStamp int64 `json:"timestamp"`
	// Nonce     int64 `json:"nonce"`
	// Iat       int64 `json:"iat"`
	// Exp       int64 `json:"exp"`
}

func (s *UserInfo) String() string {
	j, _ := json.Marshal(s)
	return string(j)
}

type MpcUserToken struct {
	UserInfo UserInfo

	jwt.RegisteredClaims
}
