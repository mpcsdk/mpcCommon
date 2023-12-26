package mq

const RiskServerMQ = "RiskServerMQ"
const RiskEngineMQ = "RiskEngineMQ"

type RiskCtrlMsg struct {
	Subject string      `json:"subject"`
	Data    interface{} `json:"data"`
}
type RiskCtrMsqRsp struct {
	Code int    `json:"code"`
	Msg  string `json:"message"`
}

const (
	RiskServerMQ_Subj_ContractRule = "ContractRule"
	RiskServerMQ_Subj_ContractAbi  = "ContractAbi"

	RiskEngineMQ_Subj_RiskCtrlRule = "RiskCtrlRule"
)

// //
// //RiskServerMQ

const (
	NoticeAdd    = "add"
	NoticeUpdate = "update"
	NoticeDelete = "delete"
)

// /
// //
type ContractNotice struct {
	// 'add' | 'update' | 'delete'
	Type            string `json:"type"`
	Id              int64  `json:"id"`
	ContractAddress string `json:"contractAddress"`
	SceneNo         string `json:"sceneNo"`
}

func (s *ContractNotice) IsValid() bool {
	if s.Type == "" || s.Id <= 0 || s.ContractAddress == "" || s.SceneNo == "" {
		return false
	}
	return true
}

// //RiskEngineMQ
const (
	//notice 验证
	NoticeVerify = "verify"
)

type RiskEngineRuleStrNotice struct {
	//up/del/verify
	Type     string `json:"type"`
	RuleName string `json:"ruleName"`
	RuleStr  string `json:"ruleStr"`
	SceneNo  string `json:"sceneNo"`
	Id       int64  `json:"id"`
}

func (s *RiskEngineRuleStrNotice) IsValid() bool {
	if s.Type == "" || s.RuleStr == "" || s.SceneNo == "" {
		return false
	}
	return true
}
