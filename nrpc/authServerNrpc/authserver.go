package authServerNrpc

import (
	"context"

	"time"

	"github.com/gogf/gf/v2/os/gctx"
	"github.com/nats-io/nats.go"
)

type NRpcServer struct {
	sub *nats.Subscription
	nc  *nats.Conn
	// cache *gcache.Cache
}

// func Init() *NrpcServer {
// 	// return Instance()
// }

//	func NewNRpcServerInstance() *NrpcServer {
//		once.Do(func() {
//			// nrpcServer = NewNRpcServer()
//		})
//		return nrpcServer
//	}
//
// /nrpc opts
type AuthRpcServerCfg struct {
	Url     string
	TimeOut int
}

func NewAuthRpcServer(ctx context.Context, cfg *AuthRpcServerCfg, authserver AuthServerServer) (*NRpcServer, error) {
	//
	nc, err := nats.Connect(cfg.Url, nats.Timeout(time.Duration(cfg.TimeOut)*time.Second))
	if err != nil {
		return nil, err
	}
	s := &NRpcServer{}
	///
	h := NewAuthServerHandler(gctx.GetInitCtx(), nc, authserver)
	sub, err := nc.QueueSubscribe(h.Subject(), "authServer", h.Handler)
	if err != nil {
		return nil, err
	}
	/////
	s.sub = sub
	s.nc = nc

	///
	return s, nil
}

// /
func WithAuthToken(fn func(context.Context, *AuthTokenReq) (*AuthTokenRes, error)) {

}
