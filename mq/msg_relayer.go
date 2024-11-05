package mq

import (
	"github.com/mpcsdk/mpcCommon/mpcdao/model/entity"
)

const Sub_RelayerFeeNotify = "RelayerFeeNotify"
const Sub_RelayerChannelNotify = "RelayerChannelNotify"

type RelayerChannelMsg struct {
	//up/del/verify
	Msg
	Data *entity.RelayerdminRelayerChannel `json:"data"`
}

func (s *RelayerChannelMsg) IsValid() bool {
	// if s.Opt == "" || s.RuleStr == "" || s.ChainId == "" {
	// 	return false
	// }
	return true
}

type RelayerFeeMsg struct {
	//up/del/verify
	Msg
	Data *entity.RelayerdminRelayerFee `json:"data"`
}
