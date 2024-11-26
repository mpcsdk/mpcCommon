package authService

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/mpcsdk/mpcCommon/authService/authServiceModel"
	"github.com/mpcsdk/mpcCommon/mpccode"
)

type MiddlewareAuthTokenInfoNrpcCfg struct {
	tokenInfoFn        func(ctx context.Context, tokenStr string) (*authServiceModel.MpcUserToken, error)
	tokenInfoErrHandle func(ctx context.Context, err error)
	middlewareFn       func(r *ghttp.Request, tokenInfo *authServiceModel.MpcUserToken)
	middlewareErrFn    func(ctx context.Context, err error)
}
type MiddlewareAuthTokenInfoNrpcOption func(*MiddlewareAuthTokenInfoNrpcCfg)

func WithTokenInfoFn(fn func(ctx context.Context, tokenStr string) (*authServiceModel.MpcUserToken, error)) MiddlewareAuthTokenInfoNrpcOption {
	return func(cfg *MiddlewareAuthTokenInfoNrpcCfg) {
		cfg.tokenInfoFn = fn
	}
}
func WithTokenInfoErrHandle(fn func(ctx context.Context, err error)) MiddlewareAuthTokenInfoNrpcOption {
	return func(cfg *MiddlewareAuthTokenInfoNrpcCfg) {
		cfg.tokenInfoErrHandle = fn
	}
}

func WithMiddlewareFn(fn func(r *ghttp.Request, tokenInfo *authServiceModel.MpcUserToken)) MiddlewareAuthTokenInfoNrpcOption {
	return func(cfg *MiddlewareAuthTokenInfoNrpcCfg) {
		cfg.middlewareFn = fn
	}
}
func WithMiddlewareErrFn(fn func(ctx context.Context, err error)) MiddlewareAuthTokenInfoNrpcOption {
	return func(cfg *MiddlewareAuthTokenInfoNrpcCfg) {
		cfg.middlewareErrFn = fn
	}
}
func BuildMiddlewareAuthTokenInfoNrpc(opts ...MiddlewareAuthTokenInfoNrpcOption) func(r *ghttp.Request) {
	s := &MiddlewareAuthTokenInfoNrpcCfg{}
	for _, opt := range opts {
		opt(s)
	}
	return func(r *ghttp.Request) {
		if s.tokenInfoFn == nil {
			r.SetParam("tokenInfo", &authServiceModel.MpcUserToken{
				UserInfo: authServiceModel.UserInfo{
					UserId: "abcd",
					AppId:  "abcd",
				},
			})
			r.Middleware.Next()
			return
		}
		tokenStr := r.Get("token").String()
		if tokenStr == "" {
			g.Log().Error(r.Context(), "TokenMiddleware tokenStr is empty")
			s.middlewareErrFn(r.Context(), mpccode.CodeParamInvalid())
		}
		tokenInfo, err := s.tokenInfoFn(r.Context(), tokenStr)
		if err != nil {
			g.Log().Error(r.Context(), "TokenMiddleware tokenInfoFn err:", err, "token:", tokenStr)
			s.tokenInfoErrHandle(r.Context(), err)
			return
		}

		r.SetParam("tokenInfo", tokenInfo)
		s.middlewareFn(r, tokenInfo)
	}
}
