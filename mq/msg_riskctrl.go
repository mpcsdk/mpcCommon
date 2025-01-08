package mq

import (
	"github.com/mpcsdk/mpcCommon/mpcdao/model/entity"
)

type RiskAdminContractMsg struct {
	//up/del/verify
	Msg
	Data *entity.RiskadminContractabi `json:"data"`
}
type RiskAdminChainMsg struct {
	//up/del/verify
	Msg
	Data *entity.RiskadminChaincfg `json:"data"`
}

// ////
// ////
type RiskAdminRiskRuleMsg struct {
	//up/del/verify
	Msg
	Data *entity.RiskadminRiskcontrolRule `json:"data"`
}
type RiskAdminRiskRuleCheckMsg struct {
	//up/del/verify
	Msg
	Data *entity.RiskadminRiskcontrolRule `json:"data"`
}

type RiskAdminRiskRuleCheckResp struct {
	Code int    `json:"code"`
	Msg  string `json:"message"`
}
