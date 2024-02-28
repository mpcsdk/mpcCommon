package mq

const Sub_ChainCfg = "ChainCgf"
const Sub_ContractAbi = "ContractCfg"
const Sub_ContractRule = "ContractRule"
const Sub_RiskRule = "RiskRule"

const Sub_RiskRuleReply = "RiskRuleReply"
const (
	OptAdd    = "add"
	OptUpdate = "update"
	OptDelete = "delete"
	OptCheck  = "check"
)

// /Sub_ChainCfg
type ChainCfgReq struct {
	Opt     string `json:"opt"`
	Id      int64  `json:"id"`
	ChainId uint64 `json:"chainId"`
	Coin    string `json:"coin"`
	Rpc     string `json:"rpc"`
}

// ContractAbi
type ContractAbiReq struct {
	Opt             string `json:"opt"`
	Id              int64  `json:"id"`
	ContractAddress string `json:"contractAddress"`
	SceneNo         string `json:"sceneNo"`
}

func (s *ContractAbiReq) IsValid() bool {
	if s.Opt == "" || s.Id <= 0 || s.ContractAddress == "" || s.SceneNo == "" {
		return false
	}
	return true
}

// /ContractRule
type ContractRuleReq struct {
	// 'add' | 'update' | 'delete'
	Opt             string `json:"opt"`
	Id              int64  `json:"id"`
	ContractAddress string `json:"contractAddress"`
	SceneNo         string `json:"sceneNo"`
}

func (s *ContractRuleReq) IsValid() bool {
	if s.Opt == "" || s.Id <= 0 || s.ContractAddress == "" || s.SceneNo == "" {
		return false
	}
	return true
}

// /RiskRule
type RiskCtrlRuleReq struct {
	//up/del/verify
	Opt      string `json:"opt"`
	Salience int    `json:"salience"`
	RuleName string `json:"ruleName"`
	RuleStr  string `json:"ruleStr"`
	SceneNo  string `json:"sceneNo"`
	Desc     string `json:"desc"`
	Id       int64  `json:"id"`
}

func (s *RiskCtrlRuleReq) IsValid() bool {
	if s.Opt == "" || s.RuleStr == "" || s.SceneNo == "" {
		return false
	}
	return true
}

// /RiskRuleReply
type RiskRuleReplyReq struct {
	Opt      string `json:"opt"`
	RuleName string `json:"ruleName"`
	RuleStr  string `json:"ruleStr"`
}
type RiskRuleReplyRsp struct {
	Code int    `json:"code"`
	Msg  string `json:"message"`
}
