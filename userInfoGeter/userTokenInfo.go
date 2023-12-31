package userInfoGeter

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/gogf/gf/v2/os/gcache"
)

type UserTokenInfoGeter struct {
	url   string
	c     *resty.Client
	cache *gcache.Cache
}
type respUserInfo struct {
	Status  int       `json:"status"`
	ErrCode int       `json:"errorCode"`
	Msg     string    `json:"msg"`
	Data    *UserInfo `json:"data"`
}
type UserInfo struct {
	Id         int    `json:"id"`
	UserId     string `json:"appPubKey"`
	Email      string `json:"email"`
	LoginType  string `json:"loginType"`
	Address    string `json:"address"`
	KeyHash    string `json:"keyHash"`
	CreateTime int64  `json:"create_time"`
}

func (s *UserTokenInfoGeter) GetUserInfo(ctx context.Context, token string) (*UserInfo, error) {
	resp, err := s.c.R().
		SetQueryParams(map[string]string{
			"token": token,
		}).
		// EnableTrace().
		Get(s.url)
	if err != nil {
		return nil, err
	}
	userInfo := respUserInfo{}
	err = json.Unmarshal(resp.Body(), &userInfo)
	if err != nil {
		err = fmt.Errorf("%+v, resp:%s", err, resp.String())
		return nil, err
	}
	return userInfo.Data, nil
}

func NewUserInfoGeter(url string) *UserTokenInfoGeter {
	c := resty.New()
	s := &UserTokenInfoGeter{
		c:   c,
		url: url,
	}
	return s
}
