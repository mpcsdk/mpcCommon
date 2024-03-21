// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Txs is the golang structure of table txs for DAO operations like Where/Data.
type Txs struct {
	g.Meta      `orm:"table:txs, do:true"`
	ChainId     interface{} //
	TxHash      interface{} //
	TxIndex     interface{} //
	BlockHash   interface{} //
	BlockNumber interface{} //
	From        interface{} //
	To          interface{} //
	Value       interface{} //
	Contract    interface{} //
	Data        interface{} //
	LogIndex    interface{} //
}
