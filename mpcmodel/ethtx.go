// EthTx is the golang structure for table eth_tx.
package mpcmodel

type EthTx struct {
	Address     string `json:"address"     ` //
	Contract    string `json:"contract"    ` //
	MethodName  string `json:"methodName"  ` //
	MethodSig   string `json:"methodSig"   ` //
	EventName   string `json:"eventName"   ` //
	EventSig    string `json:"eventSig"    ` //
	Topics      string `json:"topics"      ` //
	From        string `json:"from"        ` //
	To          string `json:"to"          ` //
	Value       string `json:"value"       ` //
	Kind        string `json:"kind"        ` //
	BlockNumber int64  `json:"blockNumber" ` //
	BlockHash   string `json:"blockHash"   ` //
	TxHash      string `json:"txHash"      ` //
	TxIndex     int64  `json:"txIndex"     ` //
	LogIndex    int64  `json:"logIndex"    ` //
	Data        string `json:"data"        ` //
	Name        string `json:"name"        ` //
}
