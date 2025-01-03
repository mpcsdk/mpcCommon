package mq

const Sub_ChainCfg = "ChainCgf"
const Sub_ContractAbi = "ContractCfg"
const Sub_ContractRule = "ContractRule"
const Sub_RiskRule = "RiskRule"

const Sub_RiskRuleReply = "RiskRuleReply"

// //
// //
const (
	OptAdd    = "add"
	OptUpdate = "update"
	OptDelete = "delete"
	OptCheck  = "check"
)

type Msg struct {
	Sub  string      `json:"sub"`
	Opt  string      `json:"opt"`
	Data interface{} `json:"data"`
}
