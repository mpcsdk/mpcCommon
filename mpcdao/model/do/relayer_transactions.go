// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// RelayerTransactions is the golang structure of table transactions for DAO operations like Where/Data.
type RelayerTransactions struct {
	g.Meta          `orm:"table:transactions, do:true"`
	RelayerIdent    interface{} //
	ChainId         interface{} // 链id
	AppId           interface{} //
	BlockNumber     interface{} //
	RelayerTxHash   interface{} // relayer根据用户交易链下计算的交易hash
	TxHash          interface{} // 链上交易的交易hash
	TransactionData interface{} // input transaction
	Target          interface{} //
	WalletAddress   interface{} // 用户钱包地址
	WalletNonce     interface{} //
	Submitter       interface{} // address of submitter
	SubmitterNonce  interface{} // the nonce of submitter
	GasLimit        interface{} // 上链时计算的gas limit
	GasPrice        interface{} // 上链时的gas price
	Discount        interface{} // 折扣百分比，值0-100
	Status          interface{} // 交易状态  @values 0 收到交易 ｜ 1 交易成功 ｜ 2 交易失败 | 3 交易丢弃
	ErrorReason     interface{} // 交易失败
	Createdat       *gtime.Time //
	Updatedat       *gtime.Time //
}
