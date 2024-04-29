// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ChainTx is the golang structure of table chain_tx for DAO operations like Where/Data.
type ChainTx struct {
	g.Meta       `orm:"table:chain_tx, do:true"`
	ChainId      interface{} //
	Height       interface{} //
	BlockHash    interface{} //
	Ts           interface{} //
	TxHash       interface{} //
	TxIdx        interface{} //
	LogIdx       interface{} //
	From         interface{} //
	To           interface{} //
	Contract     interface{} //
	Value        interface{} //
	Gas          interface{} //
	GasPrice     interface{} //
	Nonce        interface{} //
	Kind         interface{} //
	TokenId      interface{} //
	ContractName interface{} //
}
