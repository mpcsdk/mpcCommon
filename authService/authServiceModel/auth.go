package authServiceModel

import (
	"encoding/json"

	"github.com/golang-jwt/jwt/v4"
)

type UserInfo struct {
	Id         int    `json:"id"`
	UserId     string `json:"appPubKey"`
	Email      string `json:"email"`
	LoginType  string `json:"loginType"`
	Address    string `json:"address"`
	KeyHash    string `json:"keyHash"`
	CreateTime int64  `json:"create_time"`
}

func (s *UserInfo) String() string {
	j, _ := json.Marshal(s)
	return string(j)
}

type MpcUserToken struct {
	UserInfo UserInfo
	ErrCode  uint32
	ErrMsg   string
	jwt.RegisteredClaims
}
