package authServiceApi

import (
	"context"

	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/mpcsdk/mpcCommon/authServiceModel"
)

// //mpc user token
type AuthServiceApi struct {
	url     string
	timeout int64
	cache   *gcache.Cache
}

// //
func NewAuthServiceApi(r *gredis.Redis, url string, timeout int64) (*AuthServiceApi, error) {
	_, err := r.Conn(gctx.GetInitCtx())
	if err != nil {
		return nil, err
	}
	cache := gcache.New()
	cache.SetAdapter(gcache.NewAdapterRedis(r))

	return &AuthServiceApi{cache: cache, url: url, timeout: timeout}, nil
}

type respAuthToken struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Token string `json:"token"`
	} `json:"data"`
}
type respTokenInfo struct {
	Code int                            `json:"code"`
	Msg  string                         `json:"msg"`
	Data *authServiceModel.MpcUserToken `json:"data"`
}

func (s *AuthServiceApi) AuthToken(ctx context.Context, userToken string) (string, error) {
	if v, err := s.cache.Get(ctx, "AuthApi:AuthToken:"+userToken); err == nil && !v.IsEmpty() {
		return v.String(), nil
	}
	///
	res, err := g.Client().Post(ctx, s.url+"/AuthToken", g.Map{
		"userToken": userToken,
	})
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	////
	v := g.NewVar(res.ReadAll())
	var authRes *respAuthToken = nil

	err = v.Struct(&authRes)
	if err != nil {
		return "", err
	}
	//todo: expire
	s.cache.Set(ctx, "AuthApi:AuthToken:"+userToken, res, 0)
	////
	return authRes.Data.Token, nil
}

// //
func (s *AuthServiceApi) RefToken(ctx context.Context, token string) (string, error) {
	if v, err := s.cache.Get(ctx, "AuthApi:RefToken:"+token); err == nil && !v.IsEmpty() {
		return v.String(), nil
	}
	res, err := g.Client().Post(ctx, s.url+"/RefToken", g.Map{
		"token": token,
	})
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	////
	v := g.NewVar(res.ReadAll())
	var authRes *respAuthToken = nil

	err = v.Struct(&authRes)
	if err != nil {
		return "", err
	}
	//todo: expire
	s.cache.Set(ctx, "AuthApi:RefToken:"+token, res, 0)
	////
	return authRes.Data.Token, nil
}
func (s *AuthServiceApi) TokenInfo(ctx context.Context, token string) (*authServiceModel.MpcUserToken, error) {
	if v, err := s.cache.Get(ctx, "AuthApi:TokenInfo:"+token); err == nil && !v.IsEmpty() {
		var res *authServiceModel.MpcUserToken = nil
		v.Struct(&res)
		return res, nil
	}
	res, err := g.Client().Post(ctx, s.url+"/TokenInfo", g.Map{
		"token": token,
	})
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	////
	v := g.NewVar(res.ReadAll())
	var authRes *respTokenInfo = nil

	err = v.Struct(&authRes)
	if err != nil {
		return nil, err
	}
	////
	//todo: expire
	s.cache.Set(ctx, "AuthApi:TokenInfo:"+token, res, 0)
	return authRes.Data, nil
}
