package mq

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/nats-io/nats.go"
)

type NatsServer struct {
	ctx context.Context
	sub *nats.Subscription
	nc  *nats.Conn
}

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
	return s
}

func (s *NatsServer) NatsSubscribe(subj string, fn func(data []byte) error) {
	ch := make(chan *nats.Msg, 64)
	sub, err := s.nc.ChanSubscribe(subj, ch)
	if err != nil {
		panic(err)
	}
	go s.subscribe(sub, ch, fn)
}
func (s *NatsServer) NatsSubscribeReply(subj string, fn func(data []byte) ([]byte, error)) {
	////
	ch := make(chan *nats.Msg, 64)
	sub, err := s.nc.ChanQueueSubscribe(subj, subj, ch)
	if err != nil {
		panic(err)
	}
	go s.queueSubscribe(sub, ch, fn)
}

// /
func (s *NatsServer) subscribe(sub *nats.Subscription, ch chan *nats.Msg, fn func(data []byte) error) {
	for {
		select {
		case natsmsg := <-ch:
			switch natsmsg.Subject {
			case Sub_ChainCfg,
				Sub_ContractAbi,
				Sub_ContractRule,
				Sub_RiskRule:
				err := fn(natsmsg.Data)
				if err != nil {
					g.Log().Error(s.ctx, err)
					continue
				}
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