package relayerAdminServiceNats

import (
	"encoding/json"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/mpcsdk/mpcCommon/mq"
	"github.com/nats-io/nats.go"
)

// ////////////
func (s *RelayerAdminNatsService) runConsumeAppCfgFn() {
	ch := make(chan *nats.Msg, 64)
	sub, err := s.nc.ChanSubscribe(mq.Sub_RelayerAppIdNotify, ch)
	if err != nil {
		panic(err)
	}

	for {
		select {
		case msg := <-ch:
			var data mq.RelayerAdminAppCfgMsg
			var err error
			if err = json.Unmarshal(msg.Data, &data); err == nil {
				////set cfg
				s.relayerRepo.SetApp(data.Data.Id, &data)
				///call consumer
				if s.opt.consumeAppCfgFn != nil {
					err = s.opt.consumeAppCfgFn(s.ctx, &data)
				}
			}
			if err != nil {
				g.Log().Error(s.ctx, "runConsumeAppCfgFn Unmarshal:", msg.Data, ",err:", err)
			}
			msg.Ack()
		case <-s.ctx.Done():
			sub.Unsubscribe()
			close(ch)
			sub.Drain()
		}
	}
}

// //
func (s *RelayerAdminNatsService) runConsumeAssignFeeFen() {
	ch := make(chan *nats.Msg, 64)
	sub, err := s.nc.ChanSubscribe(mq.Sub_RelayerAssignFeeNotify, ch)
	if err != nil {
		panic(err)
	}

	for {
		select {
		case msg := <-ch:
			var data mq.RelayerAdminAssignFeeMsg
			var err error
			if err = json.Unmarshal(msg.Data, &data); err == nil {
				s.relayerRepo.SetAssignFee(data.Data.Id, &data)
				if s.opt.consumeAssignFeeFen != nil {
					err = s.opt.consumeAssignFeeFen(s.ctx, &data)
				}
			}
			if err != nil {
				g.Log().Error(s.ctx, "runConsumeAssignFeeFen Unmarshal:", msg.Data, ",err:", err)
			}
			msg.Ack()
		case <-s.ctx.Done():
			sub.Unsubscribe()
			close(ch)
			sub.Drain()
		}
	}
}

// //
func (s *RelayerAdminNatsService) runConsumeSpecifiedGasFn() {
	ch := make(chan *nats.Msg, 64)
	sub, err := s.nc.ChanSubscribe(mq.Sub_RelayerSpecifiedGasNotify, ch)
	if err != nil {
		panic(err)
	}
	///
	for {
		select {
		case msg := <-ch:
			var data mq.RelayerAdminSpecifiedGas
			var err error
			if err = json.Unmarshal(msg.Data, &data); err == nil {
				s.relayerRepo.SetSpecifiedGas(data.Data.Id, &data)
				if s.opt.consumeSpecifiedGasFn != nil {
					err = s.opt.consumeSpecifiedGasFn(s.ctx, &data)
				}
			}
			if err != nil {
				g.Log().Error(s.ctx, "runConsumeSpecifiedGasFn Unmarshal:", msg.Data, ",err:", err)
			}
			msg.Ack()
		case <-s.ctx.Done():
			sub.Unsubscribe()
			close(ch)
			sub.Drain()
		}
	}
}
