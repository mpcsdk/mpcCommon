package authService

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/mpcsdk/mpcCommon/authService/authServiceApi"
	"github.com/mpcsdk/mpcCommon/authService/authServiceNrpc"
)

func BuildMiddlewareParseTokenApi(authapi *authServiceApi.AuthServiceApi) func(r *ghttp.Request) {
	return func(r *ghttp.Request) {
		tokenStr := r.Get("token").String()
		mpcToken, err := authapi.TokenInfo(r.Context(), tokenStr)
		if err != nil {
			g.RequestFromCtx(r.Context()).Response.WriteStatusExit(500)
		}
		r.SetParam("mpcToken", mpcToken)
		r.Middleware.Next()
	}
}

// func BuildMiddlewareParseTokenNrpc(nrpc *authServiceNrpc.AuthClientNrpc) func(r *ghttp.Request) {
// 	return func(r *ghttp.Request) {
// 		tokenStr := r.Get("token").String()
// 		token, err := nrpc.TokenInfo(r.Context(), tokenStr)
// 		if err != nil {
// 			g.RequestFromCtx(r.Context()).Response.WriteStatusExit(500)
// 		}
// 		if !token.IsValid {
// 			g.RequestFromCtx(r.Context()).Response.WriteStatusExit(500)
// 		}

// 		// r.SetParam("mpcToken", token)
// 		r.Middleware.Next()
// 	}
// }

type MiddlewareAuthTokenInfoNrpcCfg struct {
	tokenInfoFn            func(ctx context.Context, tokenStr string) (*authServiceNrpc.TokenInfoRes, error)
	tokenInfoErrHandle     func(ctx context.Context, err error)
	tokenInfoInValidHandle func(ctx context.Context, tokenInfo *authServiceNrpc.TokenInfoRes)
	middlewareFn           func(r *ghttp.Request, tokenInfo *authServiceNrpc.TokenInfoRes)
	middlewareErrFn        func(r *ghttp.Request)
}
type MiddlewareAuthTokenInfoNrpcOption func(*MiddlewareAuthTokenInfoNrpcCfg)

func WithTokenInfoFn(fn func(ctx context.Context, tokenStr string) (*authServiceNrpc.TokenInfoRes, error)) MiddlewareAuthTokenInfoNrpcOption {
	return func(cfg *MiddlewareAuthTokenInfoNrpcCfg) {
		cfg.tokenInfoFn = fn
	}
}
func WithTokenInfoErrHandle(fn func(ctx context.Context, err error)) MiddlewareAuthTokenInfoNrpcOption {
	return func(cfg *MiddlewareAuthTokenInfoNrpcCfg) {
		cfg.tokenInfoErrHandle = fn
	}
}
func WithTokenInfoInValidHandle(fn func(ctx context.Context, tokenInfo *authServiceNrpc.TokenInfoRes)) MiddlewareAuthTokenInfoNrpcOption {
	return func(cfg *MiddlewareAuthTokenInfoNrpcCfg) {
		cfg.tokenInfoInValidHandle = fn
	}
}
func WithMiddlewareFn(fn func(r *ghttp.Request, tokenInfo *authServiceNrpc.TokenInfoRes)) MiddlewareAuthTokenInfoNrpcOption {
	return func(cfg *MiddlewareAuthTokenInfoNrpcCfg) {
		cfg.middlewareFn = fn
	}
}
func WithMiddlewareErrFn(fn func(r *ghttp.Request)) MiddlewareAuthTokenInfoNrpcOption {
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
		tokenStr := r.Get("token").String()
		if tokenStr == "" {
			s.middlewareErrFn(r)
		}
		tokenInfo, err := s.tokenInfoFn(r.Context(), tokenStr)
		if err != nil {
			s.tokenInfoErrHandle(r.Context(), err)
		}
		if !tokenInfo.IsValid {
			s.tokenInfoInValidHandle(r.Context(), tokenInfo)
		}

		r.SetParam("tokenInfo", tokenInfo)
		s.middlewareFn(r, tokenInfo)
	}
}
