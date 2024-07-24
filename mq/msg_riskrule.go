package mq

import (
	"encoding/json"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/mpcsdk/mpcCommon/mpcdao/model/entity"
	"github.com/nats-io/nats.go"
)

// /RiskRule
type RiskCtrlRuleMsg struct {
	//up/del/verify
	Msg
	Data *entity.RiskcontrolRule `json:"data"`
}

func (s *RiskCtrlRuleMsg) IsValid() bool {
	// if s.Opt == "" || s.RuleStr == "" || s.ChainId == "" {
	// 	return false
	// }
	return true
}
func (s *NatsServer) Sub_RiskCtrlRule(subj string, fn func(data *RiskCtrlRuleMsg) error) {
	ch := make(chan *nats.Msg, 64)
	sub, err := s.nc.ChanSubscribe(subj, ch)
	if err != nil {
		panic(err)
	}
	//
	go func() {
		for {
			select {
			case msg := <-ch:
				data := RiskCtrlRuleMsg{}
				var err error
				if err = json.Unmarshal(msg.Data, &data); err == nil {
					err = fn(&data)
				}
				if err != nil {
					g.Log().Error(s.ctx, "Sub_RiskCtrlRule Unmarshal:", msg.Data, ",err:", err)
				}
				msg.Ack()
			case <-s.ctx.Done():
				sub.Unsubscribe()
				close(ch)
				sub.Drain()
			}
		}
	}()
}

// /RiskRuleReply
type RiskRuleReplyMsg struct {
	Opt      string `json:"opt"`
	RuleName string `json:"ruleName"`
	RuleStr  string `json:"ruleStr"`
}
type RiskRuleReply struct {
	Code int    `json:"code"`
	Msg  string `json:"message"`
}

func (s *NatsServer) SubReply_RiskCtrlRule(subj string, fn func(data *RiskRuleReplyMsg) (*RiskRuleReply, error)) {
	ch := make(chan *nats.Msg, 64)
	sub, err := s.nc.ChanQueueSubscribe(subj, subj, ch)
	if err != nil {
		panic(err)
	}
	//
	go func() {
		for {
			select {
			case msg := <-ch:
				var data RiskRuleReplyMsg
				if err := json.Unmarshal(msg.Data, &data); err != nil {
					g.Log().Error(s.ctx, "SubReply_RiskCtrlRule Unmarshal:", msg.Data, ",err:", err)
					b, _ := json.Marshal(&RiskRuleReply{
						Code: 1,
						Msg:  err.Error(),
					})
					msg.Respond(b)
					continue
				}
				rst, err := fn(&data)
				if err != nil {
					g.Log().Error(s.ctx, "SubReply_RiskCtrlRule fn:", err)
					b, _ := json.Marshal(&RiskRuleReply{
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
	}()
}
