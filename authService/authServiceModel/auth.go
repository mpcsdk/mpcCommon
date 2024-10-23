package authServiceModel

import "github.com/golang-jwt/jwt/v4"

type UserInfo struct {
	Id         int    `json:"id"`
	UserId     string `json:"appPubKey"`
	Email      string `json:"email"`
	LoginType  string `json:"loginType"`
	Address    string `json:"address"`
	KeyHash    string `json:"keyHash"`
	CreateTime int64  `json:"create_time"`
}
type MpcUserToken struct {
	UserInfo *UserInfo
	jwt.RegisteredClaims
}
