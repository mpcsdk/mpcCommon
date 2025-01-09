// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// ChainTx is the golang structure for table chain_tx.
type ChainTx struct {
	ChainId   int64  `json:"chainId"   orm:"chain_id"   ` //
	Height    int64  `json:"height"    orm:"height"     ` //
	BlockHash string `json:"blockHash" orm:"block_hash" ` //
	Ts        int64  `json:"ts"        orm:"ts"         ` //
	TxHash    string `json:"txHash"    orm:"tx_hash"    ` //
	TxIdx     int    `json:"txIdx"     orm:"tx_idx"     ` //
	LogIdx    int    `json:"logIdx"    orm:"log_idx"    ` //
	From      string `json:"from"      orm:"from"       ` //
	To        string `json:"to"        orm:"to"         ` //
	Contract  string `json:"contract"  orm:"contract"   ` //
	Value     string `json:"value"     orm:"value"      ` //
	Gas       string `json:"gas"       orm:"gas"        ` //
	GasPrice  string `json:"gasPrice"  orm:"gas_price"  ` //
	Nonce     int64  `json:"nonce"     orm:"nonce"      ` //
	Kind      string `json:"kind"      orm:"kind"       ` //
	TokenId   string `json:"tokenId"   orm:"token_id"   ` //
	Removed   bool   `json:"removed"   orm:"removed"    ` //
	Status    int64  `json:"status"    orm:"status"     ` //
	TraceTag  string `json:"traceTag"  orm:"traceTag"   ` //
}
