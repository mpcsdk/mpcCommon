package riskAdminServiceNats

import (
	"encoding/json"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/mpcsdk/mpcCommon/mq"
	"github.com/nats-io/nats.go"
)

// ////////////
func (s *RiskAdminNatsService) runConsumeChainFn() {
	ch := make(chan *nats.Msg, 64)
	sub, err := s.nc.ChanSubscribe(mq.Sub_ChainCfg, ch)
	if err != nil {
		panic(err)
	}

	for {
		select {
		case msg := <-ch:
			var data mq.RiskAdminChainMsg
			var err error
			if err = json.Unmarshal(msg.Data, &data); err == nil {
				////set cfg
				s.riskadminCfg.SetChain(data.Data.Id, &data)
				///call consumer
				if s.opt.consumeChainFn != nil {
					err = s.opt.consumeChainFn(s.ctx, &data)
				}
			}
			if err != nil {
				g.Log().Error(s.ctx, "runConsumeChainFn Unmarshal:", msg.Data, ",err:", err)
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
func (s *RiskAdminNatsService) runConsumeContractFn() {
	ch := make(chan *nats.Msg, 64)
	sub, err := s.nc.ChanSubscribe(mq.Sub_ContractAbi, ch)
	if err != nil {
		panic(err)
	}

	for {
		select {
		case msg := <-ch:
			var data mq.RiskAdminContractMsg
			var err error
			if err = json.Unmarshal(msg.Data, &data); err == nil {
				s.riskadminCfg.SetContract(data.Data.Id, &data)
				if s.opt.consumeContractFn != nil {
					err = s.opt.consumeContractFn(s.ctx, &data)
				}
			}
			if err != nil {
				g.Log().Error(s.ctx, "runConsumeContractFn Unmarshal:", msg.Data, ",err:", err)
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
func (s *RiskAdminNatsService) runConsumeRiskCtrlRuleFn() {
	ch := make(chan *nats.Msg, 64)
	sub, err := s.nc.ChanSubscribe(mq.Sub_RiskRule, ch)
	if err != nil {
		panic(err)
	}
	///
	for {
		select {
		case msg := <-ch:
			var data mq.RiskAdminRiskRuleMsg
			var err error
			if err = json.Unmarshal(msg.Data, &data); err == nil {
				s.riskadminCfg.SetRiskRule(data.Data.Id, &data)
				if s.opt.consumeRiskRuleFn != nil {
					err = s.opt.consumeRiskRuleFn(s.ctx, &data)
				}
			}
			if err != nil {
				g.Log().Error(s.ctx, "runConsumeRiskCtrlRuleFn Unmarshal:", msg.Data, ",err:", err)
			}
			msg.Ack()
		case <-s.ctx.Done():
			sub.Unsubscribe()
			close(ch)
			sub.Drain()
		}
	}
}

// // //

func (s *RiskAdminNatsService) runConsumeRiskCtrlRuleCheckRespFn() {
	ch := make(chan *nats.Msg, 64)
	sub, err := s.nc.ChanQueueSubscribe(mq.Sub_RiskRuleReply, mq.Sub_RiskRuleReply, ch)
	if err != nil {
		panic(err)
	}
	///
	for {
		select {
		case msg := <-ch:
			var data mq.RiskAdminRiskRuleCheckMsg
			if err := json.Unmarshal(msg.Data, &data); err != nil {
				g.Log().Error(s.ctx, "runConsumeRiskCtrlRuleCheckRespFn Unmarshal:", msg.Data, ",err:", err)
				b, _ := json.Marshal(&mq.RiskAdminRiskRuleCheckResp{
					Code: 1,
					Msg:  err.Error(),
				})
				msg.Respond(b)
				continue
			}
			if s.opt.consumeRiskRuleCheckRespFn == nil {
				b, _ := json.Marshal(&mq.RiskAdminRiskRuleCheckResp{
					Code: 1,
					Msg:  "have no check function reply",
				})
				msg.Respond(b)
				continue
			}
			rst, err := s.opt.ConsumeRiskRuleCheckRespFn()(s.ctx, &data)
			if err != nil {
				g.Log().Error(s.ctx, "runConsumeRiskCtrlRuleCheckRespFn fn:", err)
				b, _ := json.Marshal(&mq.RiskAdminRiskRuleCheckResp{
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
			close(ch)
			sub.Drain()
		}
	}
}
