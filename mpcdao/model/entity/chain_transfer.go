// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// ChainTransfer is the golang structure for table chain_transfer.
type ChainTransfer struct {
	ChainId   int64  `json:"chainId"   ` //
	Height    int64  `json:"height"    ` //
	BlockHash string `json:"blockHash" ` //
	Ts        int64  `json:"ts"        ` //
	TxHash    string `json:"txHash"    ` //
	TxIdx     int    `json:"txIdx"     ` //
	LogIdx    int    `json:"logIdx"    ` //
	From      string `json:"from"      ` //
	To        string `json:"to"        ` //
	Contract  string `json:"contract"  ` //
	Value     string `json:"value"     ` //
	Gas       string `json:"gas"       ` //
	GasPrice  string `json:"gasPrice"  ` //
	Nonce     int64  `json:"nonce"     ` //
	Kind      string `json:"kind"      ` //
	TokenId   string `json:"tokenId"   ` //
	Removed   bool   `json:"removed"   ` //
	Status    int64  `json:"status"    ` //
}
