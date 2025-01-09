package relayerServiceNrpc

import (
	"time"

	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/nats-io/nats.go"
)

type RiskCtrlRpcClient struct {
	nc    *nats.Conn
	cache *gcache.Cache
}

func NewRelayerRpcClient(r *gredis.Redis, natsUrl string, timeout int64) (*RiskCtrlRpcClient, error) {
	s := &RiskCtrlRpcClient{}
	nc, err := nats.Connect(natsUrl, nats.Timeout(time.Second*time.Duration(timeout)))
	if err != nil {
		return nil, err
	}
	// cli := NewRiskCtrlServiceClient(nc)
	///
	_, err = r.Conn(gctx.GetInitCtx())
	if err != nil {
		return nil, err
	}
	cache := gcache.New()
	cache.SetAdapter(gcache.NewAdapterRedis(r))

	///
	s.nc = nc
	s.cache = cache
	return s, nil
}
