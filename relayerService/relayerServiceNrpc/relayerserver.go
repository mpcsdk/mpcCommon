package relayerServiceNrpc

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

type RelayerRpcService struct {
	ctx context.Context
	// sub               *nats.Subscription
	relayerChannelSub *nats.Subscription
	realyerFeeSub     *nats.Subscription
	nc                *nats.Conn
	// cache *gcache.Cache
	cfg *RelayerRpcServiceCfg
}

// /nrpc opts
type RelayerRpcServiceCfg struct {
	Url     string
	TimeOut int

	consumeRelayerFeeFn     func(ctx context.Context, data *mq.RelayerFeeMsg) error
	consumeRelayerChannelFn func(ctx context.Context, data *mq.RelayerChannelMsg) error
}

func RelayerRpcServiceCfgBuilder() *RelayerRpcServiceCfg {
	return &RelayerRpcServiceCfg{}
}
func (s *RelayerRpcServiceCfg) check() error {
	if s.consumeRelayerChannelFn == nil {
		return errors.New("RelayerRpcServiceCfg consumeRelayerChannelFn is nil")
	}
	if s.consumeRelayerFeeFn == nil {
		return errors.New("RelayerRpcServiceCfg consumeRelayerFeeFn is nil")
	}
	if s.Url == "" {
		return errors.New("RelayerRpcServiceCfg Url is empty")
	}
	return nil
}
func (s *RelayerRpcServiceCfg) WithConsumeRelayerChannel(consumeRelayerChannelFn func(ctx context.Context, data *mq.RelayerChannelMsg) error) *RelayerRpcServiceCfg {
	s.consumeRelayerChannelFn = consumeRelayerChannelFn
	return s
}
func (s *RelayerRpcServiceCfg) WithConsumeRelayerFee(consumeRelayerFeeFn func(ctx context.Context, data *mq.RelayerFeeMsg) error) *RelayerRpcServiceCfg {
	s.consumeRelayerFeeFn = consumeRelayerFeeFn
	return s
}
func (s *RelayerRpcServiceCfg) WithUrl(url string) *RelayerRpcServiceCfg {
	s.Url = url
	return s
}
func (s *RelayerRpcServiceCfg) WithTimeOut(TimeOut int) *RelayerRpcServiceCfg {
	s.TimeOut = TimeOut
	return s
}

// /////
// /////
func NewRelayerRpcService(ctx context.Context, cfg *RelayerRpcServiceCfg) (*RelayerRpcService, error) {
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
	s := &RelayerRpcService{
		ctx: gctx.GetInitCtx(),
		cfg: cfg,
	}
	///
	// h := NewRiskCtrlServiceHandler(gctx.GetInitCtx(), nc, server)
	// sub, err := nc.QueueSubscribe(h.Subject(), h.Subject(), h.Handler)
	// if err != nil {
	// 	return nil, err
	// }
	/////
	// s.sub = sub
	s.nc = nc
	///relayerchannel
	chRelayerChannel := make(chan *nats.Msg, 64)
	replyChannelSub, err := nc.ChanQueueSubscribe(mq.Sub_RelayerChannelNotify, mq.Sub_RelayerChannelNotify, chRelayerChannel)
	if err != nil {
		panic(err)
	}
	s.relayerChannelSub = replyChannelSub
	////
	go func() {
		for {
			select {
			case msg := <-chRelayerChannel:
				var data mq.RelayerChannelMsg
				var err error
				if err = json.Unmarshal(msg.Data, &data); err == nil {
					err = s.cfg.consumeRelayerChannelFn(s.ctx, &data)
				}
				if err != nil {
					g.Log().Error(s.ctx, "Sub_RiskCtrlRule Unmarshal:", msg.Data, ",err:", err)
				}
				msg.Ack()
			case <-s.ctx.Done():
				replyChannelSub.Unsubscribe()
				close(chRelayerChannel)
				replyChannelSub.Drain()
			}
		}
	}()
	////relayerfee
	chRelayerFee := make(chan *nats.Msg, 64)
	realyerFeeSub, err := nc.ChanQueueSubscribe(mq.Sub_RelayerFeeNotify, mq.Sub_RelayerFeeNotify, chRelayerFee)
	if err != nil {
		panic(err)
	}
	s.realyerFeeSub = realyerFeeSub
	////
	go func() {
		for {
			select {
			case msg := <-chRelayerFee:
				data := mq.RelayerFeeMsg{}
				var err error
				if err = json.Unmarshal(msg.Data, &data); err == nil {
					err = s.cfg.consumeRelayerFeeFn(s.ctx, &data)
				}
				if err != nil {
					g.Log().Error(s.ctx, "Sub_RiskCtrlRule Unmarshal:", msg.Data, ",err:", err)
				}
				msg.Ack()
			case <-s.ctx.Done():
				realyerFeeSub.Unsubscribe()
				close(chRelayerFee)
				realyerFeeSub.Drain()
			}
		}
	}()
	return s, nil
}

// //
