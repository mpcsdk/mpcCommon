package riskAdminServiceNats

import (
	"time"

	"github.com/franklihub/nrpc"
	"github.com/nats-io/nats.go"
)

type RiskAdminNatsClient struct {
	nc *nats.Conn
	// cache *gcache.Cache
	///
}

func NewRiskAdminNatsClient(natsUrl string, timeout int) (*RiskAdminNatsClient, error) {
	s := &RiskAdminNatsClient{}
	nc, err := nats.Connect(natsUrl, nats.Timeout(time.Second*time.Duration(timeout)))
	if err != nil {
		return nil, err
	}
	///
	// _, err = r.Conn(gctx.GetInitCtx())
	// if err != nil {
	// 	return nil, err
	// }
	// cache := gcache.New()
	// cache.SetAdapter(gcache.NewAdapterRedis(r))
	// s.cache = cache

	///
	s.nc = nc
	return s, nil
}
func (s *RiskAdminNatsClient) Flush() {
	err := s.nc.Flush()
	if err != nil {
		panic(err)
	}
}

func (s *RiskAdminNatsClient) TryFlush(err error) {
	if _, ok := err.(*nrpc.Error); ok {
		return
	} else {
		if err == nats.ErrTimeout {

		} else {
			return

		}
	}
	err = s.nc.Flush()
	if err != nil {
		panic(err)
	}
}

func (s *RiskAdminNatsClient) TestSendMsg(sub string, data []byte) {
	err := s.nc.Publish(sub, data)
	if err != nil {
		panic(err)
	}
}
func (s *RiskAdminNatsClient) TestSendReplyMsg(sub string, data []byte) []byte {
	replyMsg, err := s.nc.Request(sub, data, 10*time.Second)
	if err != nil {
		panic(err)
	}
	return replyMsg.Data
}
