package authServiceNrpc

import (
	"context"

	"time"

	"github.com/gogf/gf/v2/os/gctx"
	"github.com/nats-io/nats.go"
)

type AuthRpcService struct {
	sub *nats.Subscription
	nc  *nats.Conn
	// cache *gcache.Cache
}

// /nrpc opts
type AuthRpcServiceCfg struct {
	Url     string
	TimeOut int64
}

func NewAuthRpcService(ctx context.Context, cfg *AuthRpcServiceCfg, authserver AuthServiceServer) (*AuthRpcService, error) {
	//
	nc, err := nats.Connect(cfg.Url, nats.Timeout(time.Duration(cfg.TimeOut)*time.Second))
	if err != nil {
		return nil, err
	}
	s := &AuthRpcService{}
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
