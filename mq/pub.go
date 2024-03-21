package mq

const Sub_ChainCfg = "ChainCgf"
const Sub_ContractAbi = "ContractCfg"
const Sub_ContractRule = "ContractRule"
const Sub_RiskRule = "RiskRule"
const Sub_ChainTx = "ChainTx"

const Sub_RiskRuleReply = "RiskRuleReply"
const (
	OptAdd    = "add"
	OptUpdate = "update"
	OptDelete = "delete"
	OptCheck  = "check"
)

// Sub_ChainTx
type ChainTxMsg struct {
	ChainId     int64  `json:"chainId"`
	TxHash      string `json:"txHash"`
	TxIdx       int    `json:"txIdx"`
	BlockNumber int    `json:"blockNumber"`
	BlockHash   string `json:"blockHash"`
	LogIdx      int    `json:"logIdx"`

	From     string `json:"from"`
	To       string `json:"to"`
	Contract string `json:"contract"`
	Value    string `json:"value"`
}

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
	ChainId         string `json:"chainId"`
}

func (s *ContractAbiReq) IsValid() bool {
	if s.Opt == "" || s.Id <= 0 || s.ContractAddress == "" || s.ChainId == "" {
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
	ChainId         string `json:"chainId"`
}

func (s *ContractRuleReq) IsValid() bool {
	if s.Opt == "" || s.Id <= 0 || s.ContractAddress == "" || s.ChainId == "" {
		return false
	}
	return true
}

// /RiskRule
type RiskCtrlRuleReq struct {
	//up/del/verify
	Opt string `json:"opt"`
	// Salience int    `json:"salience"`
	// RuleName string `json:"ruleName"`
	// RuleStr  string `json:"ruleStr"`
	// Desc     string `json:"desc"`
	Id       int  `json:"id"`
	IsEnable bool `json:"isEnable"`
}

func (s *RiskCtrlRuleReq) IsValid() bool {
	// if s.Opt == "" || s.RuleStr == "" || s.ChainId == "" {
	// 	return false
	// }
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
