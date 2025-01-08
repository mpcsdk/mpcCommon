package mq

const Sub_ChainCfg = "ChainCfg"
const Sub_ContractAbi = "ContractCfg"
const Sub_ContractRule = "ContractRule"
const Sub_RiskRule = "RiskRule"

const Sub_RiskRuleReply = "RiskRuleReply"

// //
// //
type MsgOpt string

const (
	OptAdd    MsgOpt = "add"
	OptUpdate        = "update"
	OptDelete        = "delete"
	OptCheck         = "check"
)

type Msg struct {
	Sub  string      `json:"sub"`
	Opt  MsgOpt      `json:"opt"`
	Data interface{} `json:"data"`
}
