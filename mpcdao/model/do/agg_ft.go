// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// AggFt is the golang structure of table aggFt for DAO operations like Where/Data.
type AggFt struct {
	g.Meta     `orm:"table:aggFt, do:true"`
	TxHash     interface{} //
	FromAddr   interface{} //
	ToAddr     interface{} //
	Contract   interface{} //
	Value      interface{} //
	StartTs    interface{} //
	ChainId    interface{} //
	EndTs      interface{} //
	StartBlock interface{} //
	EndBlock   interface{} //
}
