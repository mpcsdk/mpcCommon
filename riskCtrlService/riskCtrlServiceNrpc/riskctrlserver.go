package riskCtrlServiceNrpc

import (
	"context"
	"encoding/json"
	"errors"

	"time"

	"github.com/gogf/gf/v2/frame/g"
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

	checkRiskRulelFn   func(*mq.RiskRuleReplyMsg) (*mq.RiskRuleReply, error)
	consumerRiskCtrlFn func(ctx context.Context, data *mq.RiskCtrlRuleMsg) error
}

func RiskCtrlRpcServiceCfgBuilder() *RiskCtrlRpcServiceCfg {
	return &RiskCtrlRpcServiceCfg{}
}
func (s *RiskCtrlRpcServiceCfg) check() error {
	if s.checkRiskRulelFn == nil {
		return errors.New("RiskCtrlRpcServiceCfg checkRiskRulelFn is nil")
	}
	if s.consumerRiskCtrlFn == nil {
		return errors.New("RiskCtrlRpcServiceCfg consumerRiskCtrlFn is nil")
	}
	if s.Url == "" {
		return errors.New("RiskCtrlRpcServiceCfg Url is empty")
	}
	return nil
}
func (s *RiskCtrlRpcServiceCfg) WithConsumeRiskRule(consumerRiskCtrlFn func(ctx context.Context, data *mq.RiskCtrlRuleMsg) error) *RiskCtrlRpcServiceCfg {
	s.consumerRiskCtrlFn = consumerRiskCtrlFn
	return s
}
func (s *RiskCtrlRpcServiceCfg) WithCheckRiskRule(checkRiskRulelFn func(data *mq.RiskRuleReplyMsg) (*mq.RiskRuleReply, error)) *RiskCtrlRpcServiceCfg {
	s.checkRiskRulelFn = checkRiskRulelFn
	return s
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
	go func() {
		for {
			select {
			case msg := <-chReplySub:
				var data mq.RiskRuleReplyMsg
				if err := json.Unmarshal(msg.Data, &data); err != nil {
					g.Log().Error(s.ctx, "SubReply_RiskCtrlRule Unmarshal:", msg.Data, ",err:", err)
					b, _ := json.Marshal(&mq.RiskRuleReply{
						Code: 1,
						Msg:  err.Error(),
					})
					msg.Respond(b)
					continue
				}
				rst, err := cfg.checkRiskRulelFn(&data)
				if err != nil {
					g.Log().Error(s.ctx, "SubReply_RiskCtrlRule fn:", err)
					b, _ := json.Marshal(&mq.RiskRuleReply{
						Code: 1,
						Msg:  err.Error(),
					})
					msg.Respond(b)
					continue
				}
				b, _ := json.Marshal(rst)
				msg.Respond(b)
			case <-s.ctx.Done():
				sub.Unsubscribe()
				close(chReplySub)
				sub.Drain()
			}
		}
	}()
	////
	chConsumeSub := make(chan *nats.Msg, 64)
	consumerSub, err := nc.ChanQueueSubscribe(mq.Sub_RiskRuleReply, mq.Sub_RiskRuleReply, chConsumeSub)
	if err != nil {
		panic(err)
	}
	s.consumerSub = consumerSub
	////
	go func() {
		for {
			select {
			case msg := <-chConsumeSub:
				data := mq.RiskCtrlRuleMsg{}
				var err error
				if err = json.Unmarshal(msg.Data, &data); err == nil {
					err = s.cfg.consumerRiskCtrlFn(s.ctx, &data)
				}
				if err != nil {
					g.Log().Error(s.ctx, "Sub_RiskCtrlRule Unmarshal:", msg.Data, ",err:", err)
				}
				msg.Ack()
			case <-s.ctx.Done():
				sub.Unsubscribe()
				close(chConsumeSub)
				sub.Drain()
			}
		}
	}()
	return s, nil
}

// //
