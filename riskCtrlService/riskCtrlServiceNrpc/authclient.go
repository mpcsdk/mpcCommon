package riskCtrlServiceNrpc

import (
	"context"
	"errors"
	"time"

	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/types/known/emptypb"
)

type RiskCtrlRpcClient struct {
	cli   *RiskCtrlServiceClient
	nc    *nats.Conn
	cache *gcache.Cache
}

func NewRiskCtrlRpcClient(r *gredis.Redis, natsUrl string, timeout int64) (*RiskCtrlRpcClient, error) {
	s := &RiskCtrlRpcClient{}
	nc, err := nats.Connect(natsUrl, nats.Timeout(time.Second*time.Duration(timeout)))
	if err != nil {
		return nil, err
	}
	cli := NewRiskCtrlServiceClient(nc)
	///
	_, err = r.Conn(gctx.GetInitCtx())
	if err != nil {
		return nil, err
	}
	cache := gcache.New()
	cache.SetAdapter(gcache.NewAdapterRedis(r))

	///
	s.nc = nc
	s.cli = cli
	s.cache = cache
	return s, nil
}
func (s *RiskCtrlRpcClient) Flush() {
	err := s.nc.Flush()
	if err != nil {
		panic(err)
	}
	s.cli = NewRiskCtrlServiceClient(s.nc)
}

// ///
var errDeadLine = errors.New("nats: timeout")

// ///
func (s *RiskCtrlRpcClient) Alive(ctx context.Context) error {
	_, err := s.cli.Alive(&emptypb.Empty{})
	if err != nil {
		if err.Error() == errDeadLine.Error() {
			g.Log().Warning(ctx, "Alive TimeOut:")
			s.Flush()
			return nil
		}
	}
	return err
}
