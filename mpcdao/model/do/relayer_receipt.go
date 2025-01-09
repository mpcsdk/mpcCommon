// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// RelayerReceipt is the golang structure of table receipt for DAO operations like Where/Data.
type RelayerReceipt struct {
	g.Meta            `orm:"table:receipt, do:true"`
	RelayerIdent      interface{} //
	RelayerTxHash     interface{} //
	ChainId           interface{} //
	BlockHash         interface{} //
	BlockNumber       interface{} //
	TxIdx             interface{} //
	TxHash            interface{} //
	CumulativeGasUsed interface{} //
	GasUsed           interface{} //
	EffectiveGasPrice interface{} //
	ContractAddress   interface{} //
	Type              interface{} //
	Root              interface{} //
	Bloom             interface{} //
	Logs              interface{} //
	Status            interface{} //
	TransactionFee    interface{} //
	Createdat         *gtime.Time //
	AppId             interface{} //
	UserId            interface{} //
	Payer             interface{} //
}
