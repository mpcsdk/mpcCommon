package riskAdminServiceNats

import (
	"encoding/json"
	"testing"

	"github.com/mpcsdk/mpcCommon/mpcdao/model/entity"
	"github.com/mpcsdk/mpcCommon/mq"
)

func TestSend_chain(t *testing.T) {
	cli, err := NewRiskAdminNatsClient("nats://127.0.0.1:4222", 10)
	if err != nil {
		panic(err)
	}
	msg := mq.RiskAdminChainMsg{
		Msg: mq.Msg{
			Sub: mq.Sub_ChainCfg,
			Opt: mq.OptAdd,
		},
		Data: &entity.RiskadminChaincfg{
			Id:       2,
			ChainId:  95271,
			IsEnable: 1,
		},
	}
	data, _ := json.Marshal(msg)
	cli.TestSendMsg(mq.Sub_ChainCfg, data)
}
