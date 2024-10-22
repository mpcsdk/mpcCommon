package authServiceNrpc

import (
	"context"

	"time"

	"github.com/gogf/gf/v2/os/gctx"
	"github.com/nats-io/nats.go"
)

type AuthServiceNrpc struct {
	sub *nats.Subscription
	nc  *nats.Conn
	// cache *gcache.Cache
}

// /nrpc opts
type AuthServiceNrpcCfg struct {
	Url     string
	TimeOut int64
}

func NewAuthServiceNrpc(ctx context.Context, cfg *AuthServiceNrpcCfg, authserver AuthServiceServer) (*AuthServiceNrpc, error) {
	//
	nc, err := nats.Connect(cfg.Url, nats.Timeout(time.Duration(cfg.TimeOut)*time.Second))
	if err != nil {
		return nil, err
	}
	s := &AuthServiceNrpc{}
	///
	h := NewAuthServiceHandler(gctx.GetInitCtx(), nc, authserver)
	sub, err := nc.QueueSubscribe(h.Subject(), "AuthServer", h.Handler)
	if err != nil {
		return nil, err
	}
	/////
	s.sub = sub
	s.nc = nc

	///
	return s, nil
}

// //
func WithAuthToken(fn func(context.Context, *AuthTokenReq) (*AuthTokenRes, error)) {

}
