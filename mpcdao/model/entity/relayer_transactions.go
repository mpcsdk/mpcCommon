// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// RelayerTransactions is the golang structure for table relayer_transactions.
type RelayerTransactions struct {
	RelayerIdent    string      `json:"relayerIdent"    orm:"relayer_ident"    ` //
	ChainId         int64       `json:"chainId"         orm:"chain_id"         ` // 链id
	AppId           string      `json:"appId"           orm:"app_id"           ` //
	BlockNumber     int64       `json:"blockNumber"     orm:"block_number"     ` //
	RelayerTxHash   string      `json:"relayerTxHash"   orm:"relayer_tx_hash"  ` // relayer根据用户交易链下计算的交易hash
	TxHash          string      `json:"txHash"          orm:"tx_hash"          ` // 链上交易的交易hash
	TransactionData string      `json:"transactionData" orm:"transaction_data" ` // input transaction
	Target          string      `json:"target"          orm:"target"           ` //
	WalletAddress   string      `json:"walletAddress"   orm:"wallet_address"   ` // 用户钱包地址
	WalletNonce     int64       `json:"walletNonce"     orm:"wallet_nonce"     ` //
	Submitter       string      `json:"submitter"       orm:"submitter"        ` // address of submitter
	SubmitterNonce  int64       `json:"submitterNonce"  orm:"submitter_nonce"  ` // the nonce of submitter
	GasLimit        int64       `json:"gasLimit"        orm:"gas_limit"        ` // 上链时计算的gas limit
	GasPrice        string      `json:"gasPrice"        orm:"gas_price"        ` // 上链时的gas price
	Discount        int         `json:"discount"        orm:"discount"         ` // 折扣百分比，值0-100
	Status          int         `json:"status"          orm:"status"           ` // 交易状态  @values 0 收到交易 ｜ 1 交易成功 ｜ 2 交易失败 | 3 交易丢弃
	ErrorReason     string      `json:"errorReason"     orm:"error_reason"     ` // 交易失败
	Createdat       *gtime.Time `json:"createdat"       orm:"createdat"        ` //
	Updatedat       *gtime.Time `json:"updatedat"       orm:"updatedat"        ` //
	RealyerFee string `json:"realyerFee"       orm:"realyer_fee"        `
}
