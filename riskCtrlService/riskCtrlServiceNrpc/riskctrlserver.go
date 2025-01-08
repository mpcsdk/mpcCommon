package riskCtrlServiceNrpc

import (
	"context"
	"errors"

	"time"

	"github.com/gogf/gf/v2/os/gctx"
	"github.com/mpcsdk/mpcCommon/mq"
	"github.com/nats-io/nats.go"
)

type RiskCtrlRpcService struct {
	ctx         context.Context
	sub         *nats.Subscription
	replySub    *nats.Subscription
	consumerSub *nats.Subscription
	nc          *nats.Conn
	// cache *gcache.Cache
	cfg *RiskCtrlRpcServiceCfg
}

// /nrpc opts
type RiskCtrlRpcServiceCfg struct {
	Url     string
	TimeOut int
}

func RiskCtrlRpcServiceCfgBuilder() *RiskCtrlRpcServiceCfg {
	return &RiskCtrlRpcServiceCfg{}
}
func (s *RiskCtrlRpcServiceCfg) check() error {

	if s.Url == "" {
		return errors.New("RiskCtrlRpcServiceCfg Url is empty")
	}
	return nil
}

func (s *RiskCtrlRpcServiceCfg) WithUrl(url string) *RiskCtrlRpcServiceCfg {
	s.Url = url
	return s
}
func (s *RiskCtrlRpcServiceCfg) WithTimeOut(TimeOut int) *RiskCtrlRpcServiceCfg {
	s.TimeOut = TimeOut
	return s
}
func NewRiskCtrlRpcService(ctx context.Context, cfg *RiskCtrlRpcServiceCfg, server RiskCtrlServiceServer) (*RiskCtrlRpcService, error) {
	//
	err := cfg.check()
	if err != nil {
		panic(err)
	}
	//
	nc, err := nats.Connect(cfg.Url, nats.Timeout(time.Duration(cfg.TimeOut)*time.Second))
	if err != nil {
		return nil, err
	}
	s := &RiskCtrlRpcService{
		ctx: gctx.GetInitCtx(),
		cfg: cfg,
	}
	///
	h := NewRiskCtrlServiceHandler(gctx.GetInitCtx(), nc, server)
	sub, err := nc.QueueSubscribe(h.Subject(), h.Subject(), h.Handler)
	if err != nil {
		return nil, err
	}
	/////
	s.sub = sub
	s.nc = nc
	///
	chReplySub := make(chan *nats.Msg, 64)
	replySub, err := nc.ChanQueueSubscribe(mq.Sub_RiskRuleReply, mq.Sub_RiskRuleReply, chReplySub)
	if err != nil {
		panic(err)
	}
	s.replySub = replySub
	////

	////
	chConsumeSub := make(chan *nats.Msg, 64)
	consumerSub, err := nc.ChanQueueSubscribe(mq.Sub_RiskRuleReply, mq.Sub_RiskRuleReply, chConsumeSub)
	if err != nil {
		panic(err)
	}
	s.consumerSub = consumerSub
	////

	return s, nil
}

// //
