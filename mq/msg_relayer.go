package mq

import (
	"github.com/mpcsdk/mpcCommon/mpcdao/model/entity"
)

const Sub_RelayerFeeNotify = "RelayerFeeNotify"
const Sub_RelayerSpecifiedGasNotify = "RelayerSpecifiedGasNotify"
const Sub_RelayerAssignFeeNotify = "RelayerAssignFee"

type RelayerAdminAppCfgMsg struct {
	//up/del/verify
	Msg
	Data *entity.RelayeradminAppCfg `json:"data"`
}

type RelayerAdminAssignFeeMsg struct {
	//up/del/verify
	Msg
	Data *entity.RelayeradminAssignFee `json:"data"`
}
type RelayerAdminSpecifiedGas struct {
	//up/del/verify
	Msg
	Data *entity.RelayeradminSpecifiedGas `json:"data"`
}
