package mq

import (
	"github.com/mpcsdk/mpcCommon/mpcdao/model/entity"
)

type RelayerChannelMsg struct {
	//up/del/verify
	Msg
	Data *entity.RelayerChannel `json:"data"`
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
	Data *entity.RelayerFee `json:"data"`
}
