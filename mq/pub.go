package mq

type RiskCtrlKind string

const (
	RiskCtrlKind_ContractRule = "ContractRule"
	RiskCtrlKind_ContractAbi  = "ContractAbi"
)

type RiskCtrlMQ struct {
	//ContractRule
	Kind string      `json:"kind"`
	Data interface{} `json:"data"`
}

func (s *RiskCtrlMQ) GetKind() RiskCtrlKind {
	return RiskCtrlKind(s.Kind)
}

type NoticeKind string

const (
	NoticeAdd    = "add"
	NoticeUpdate = "update"
	NoticeDelete = "delete"
)

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
func (s *ContractNotice) GetType() NoticeKind {
	return NoticeKind(s.Type)
}
