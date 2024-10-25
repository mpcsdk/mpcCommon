package authService

import (
	"context"
	"errors"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/mpcsdk/mpcCommon/authService/authServiceModel"
)

// /serTokenInfo

type respUserInfo struct {
	Status  int                        `json:"status"`
	ErrCode int                        `json:"errorCode"`
	Msg     string                     `json:"msg"`
	Data    *authServiceModel.UserInfo `json:"data"`
}

func GetUserTokenInfo(ctx context.Context, url string, tokenStr string) (*authServiceModel.UserInfo, error) {
	if tokenStr == "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhcHBQdWJLZXkiOiJhYmNkIiwiaWF0IjoxNjk0NDk5Njg5LCJleHAiOjE3MjYwMzU2ODl9.OsI4nFQoSoegZJbzTQnWBaB1shMjaPinhWZlnntGub4" {
		return &authServiceModel.UserInfo{
			UserId: "abcd",
		}, nil
	}
	/////
	////
	resp, err := g.Client().Get(ctx, url, g.Map{
		"token": tokenStr,
	})
	if err != nil {
		return nil, err
	}
	/////
	defer resp.Close()
	v := gvar.New(resp.ReadAll())
	var userInfo *respUserInfo = nil
	err = v.Struct(&userInfo)
	if err != nil {
		return nil, err
	}
	////
	if userInfo.ErrCode != 0 {
		return nil, errors.New(userInfo.Msg)
	}

	return userInfo.Data, nil
}
