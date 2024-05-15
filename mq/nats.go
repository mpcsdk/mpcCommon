package mq

import (
	"context"
	"sync"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

type NatsServer struct {
	ctx context.Context
	nc  *nats.Conn
	///
	jets jetstream.JetStream
}

// /
var once sync.Once
var instance *NatsServer = nil

// /
func New(urls string) *NatsServer {
	//
	nc, err := nats.Connect(urls, nats.Timeout(5*time.Second))
	if err != nil {
		panic(err)
	}
	s := &NatsServer{}
	s.nc = nc
	s.ctx = gctx.GetInitCtx()
	///
	jets, err := jetstream.New(s.nc)
	if err != nil {
		panic(err)
	}
	//
	s.jets = jets
	///
	return s
}
func (s *NatsServer) Conn() *nats.Conn {
	return s.nc
}

// /
func (s *NatsServer) subscribe(sub *nats.Subscription, ch chan *nats.Msg, fn func(data []byte) error) {
	for {
		select {
		case natsmsg := <-ch:
			err := fn(natsmsg.Data)
			if err != nil {
				g.Log().Error(s.ctx, err)
				continue
			}
			//
		case <-s.ctx.Done():
			sub.Unsubscribe()
			close(ch)
			sub.Drain()
		}
	}
}
func (s *NatsServer) queueSubscribe(sub *nats.Subscription, ch chan *nats.Msg, fn func(data []byte) ([]byte, error)) {
	for {
		select {
		case natsmsg := <-ch:
			switch natsmsg.Subject {
			case Sub_RiskRuleReply:
				b, err := fn(natsmsg.Data)
				if err != nil {
					g.Log().Error(s.ctx, err)
					continue
				} ///
				natsmsg.Respond(b)
			}

		case <-s.ctx.Done():
			sub.Unsubscribe()
			close(ch)
			sub.Drain()
		}
	}
}
