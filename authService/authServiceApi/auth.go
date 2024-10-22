package authServiceApi

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
)

// //mpc user token
type AuthServiceApi struct {
	url     string
	timeout int64
}

// //
func NewAuthServiceApi(url string, timeout int64) *AuthServiceApi {
	return &AuthServiceApi{url: url, timeout: timeout}
}

type respAuthToken struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Token string `json:"token"`
	} `json:"data"`
}

func (s *AuthServiceApi) AuthToken(ctx context.Context, userToken string) (string, error) {

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
	////
	return authRes.Data.Token, nil
}

// //
func (s *AuthServiceApi) RefToken(ctx context.Context, token string) (string, error) {

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
	////
	return authRes.Data.Token, nil
}
func (s *AuthServiceApi) TokenInfo(ctx context.Context, token string) (string, error) {

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
	////
	return authRes.Data.Token, nil
}
