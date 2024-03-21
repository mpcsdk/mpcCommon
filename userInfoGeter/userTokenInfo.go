package userInfoGeter

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/mpcsdk/mpcCommon/mpccode"
)

// /
type UserTokenInfoGeter struct {
	url   string
	cli   *resty.Client
	cache *gcache.Cache
	dur   time.Duration
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

func (s *UserInfo) String() string {
	d, _ := json.Marshal(s)
	return string(d)
}

// /
func (s *UserTokenInfoGeter) getUserInfoCache(ctx context.Context, userToken string) (*UserInfo, error) {
	if userToken == "" {
		return nil, mpccode.CodeParamInvalid()
	}
	///
	if v, err := s.cache.Get(ctx, userToken); err == nil && !v.IsEmpty() {
		info := &UserInfo{}
		err = v.Struct(info)
		if err != nil {
			return nil, mpccode.CodeInternalError()
		}
		return info, nil
	}
	return nil, nil
}
func (s *UserTokenInfoGeter) setCache(ctx context.Context, userToken string, info *UserInfo) {
	s.cache.Set(ctx, userToken, info, 0)
}
func (s *UserTokenInfoGeter) GetUserInfo(ctx context.Context, token string) (*UserInfo, error) {
	////
	info, err := s.getUserInfoCache(ctx, token)
	if info != nil {
		return info, nil
	}
	////
	resp, err := s.cli.R().
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
	if userInfo.Status != 1 {
		return nil, mpccode.CodeTokenInvalid()
	}
	///
	s.setCache(ctx, token, userInfo.Data)
	///
	return userInfo.Data, nil
}

func NewUserInfoGeter(url string, cache *gcache.Cache, dur time.Duration) *UserTokenInfoGeter {
	c := resty.New()
	s := &UserTokenInfoGeter{
		cli:   c,
		url:   url,
		cache: cache,
		dur:   dur,
	}
	return s
}
