// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// RelayerReceipt is the golang structure for table relayer_receipt.
type RelayerReceipt struct {
	RelayerIdent      string      `json:"relayerIdent"      orm:"relayer_ident"       ` //
	RelayerTxHash     string      `json:"relayerTxHash"     orm:"relayer_tx_hash"     ` //
	ChainId           int64       `json:"chainId"           orm:"chain_id"            ` //
	BlockHash         string      `json:"blockHash"         orm:"block_hash"          ` //
	BlockNumber       int64       `json:"blockNumber"       orm:"block_number"        ` //
	TxIdx             int64       `json:"txIdx"             orm:"tx_idx"              ` //
	TxHash            string      `json:"txHash"            orm:"tx_hash"             ` //
	CumulativeGasUsed int64       `json:"cumulativeGasUsed" orm:"cumulative_gas_used" ` //
	GasUsed           int64       `json:"gasUsed"           orm:"gas_used"            ` //
	EffectiveGasPrice int64       `json:"effectiveGasPrice" orm:"effective_gas_price" ` //
	ContractAddress   string      `json:"contractAddress"   orm:"contract_address"    ` //
	Type              int64       `json:"type"              orm:"type"                ` //
	Root              string      `json:"root"              orm:"root"                ` //
	Bloom             string      `json:"bloom"             orm:"bloom"               ` //
	Logs              string      `json:"logs"              orm:"logs"                ` //
	Status            int64       `json:"status"            orm:"status"              ` //
	TransactionFee    float64     `json:"transactionFee"    orm:"transaction_fee"     ` //
	Createdat         *gtime.Time `json:"createdat"         orm:"createdat"           ` //
	AppId             string      `json:"appId"             orm:"app_id"              ` //
	UserId            string      `json:"userId"            orm:"user_id"             ` //
	Payer             string      `json:"payer"             orm:"payer"               ` //
}
