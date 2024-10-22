package authServiceModel

import (
	"encoding/json"

	"github.com/golang-jwt/jwt/v4"
)

type MpcUserToken struct {
	UserInfo *UserInfo `json:"userInfo"`
	jwt.RegisteredClaims
}
type UserInfo struct {
	Id         int    `json:"id"`
	AppId      int    `json:"appId"`
	UserId     string `json:"appPubKey"`
	Email      string `json:"email"`
	LoginType  string `json:"loginType"`
	Address    string `json:"address"`
	KeyHash    string `json:"keyHash"`
	CreateTime int64  `json:"create_time"`
}

func (s *UserInfo) String() string {
	d, _ := json.Marshal(s)
	return string(d)
}
