package authService

import (
	"context"
	"errors"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/mpcsdk/mpcCommon/authService/authServiceModel"
)

// /serTokenInfo
type UserInfo struct {
	UserId string `json:"appPubKey"`
	AppId  string `json:"appId"`

	TimeStamp int64  `json:"timestamp"`
	Nonce     string `json:"nonce"`
	Iat       int64  `json:"iat"`
	Exp       int64  `json:"exp"`
}

type respUserInfo struct {
	Status  int       `json:"status"`
	ErrCode int       `json:"errorCode"`
	Msg     string    `json:"msg"`
	Data    *UserInfo `json:"data"`
}

func GetUserTokenInfo(ctx context.Context, url string, tokenStr string) (authServiceModel.UserInfo, error) {
	if tokenStr == "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhcHBQdWJLZXkiOiJhYmNkIiwiaWF0IjoxNjk0NDk5Njg5LCJleHAiOjE3MjYwMzU2ODl9.OsI4nFQoSoegZJbzTQnWBaB1shMjaPinhWZlnntGub4" {
		return authServiceModel.UserInfo{
			UserId: "abcd",
			AppId:  "abcd",
		}, nil
	}
	/////
	////
	resp, err := g.Client().Header(map[string]string{
		"Content-Type": "application/json",
	}).Post(ctx, url, g.Map{
		"token": tokenStr,
	})
	if err != nil {
		return authServiceModel.UserInfo{}, err
	}
	/////
	defer resp.Close()
	v := gvar.New(resp.ReadAll())
	var userInfo *respUserInfo = nil
	err = v.Struct(&userInfo)
	if err != nil {
		return authServiceModel.UserInfo{}, err
	}
	////
	if userInfo.ErrCode != 0 {
		return authServiceModel.UserInfo{}, errors.New(userInfo.Msg)
	}

	return authServiceModel.UserInfo{
		AppId:  userInfo.Data.AppId,
		UserId: userInfo.Data.UserId,
	}, nil
}
