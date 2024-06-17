// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// WalletAddr is the golang structure of table wallet_addr for DAO operations like Where/Data.
type WalletAddr struct {
	g.Meta     `orm:"table:wallet_addr, do:true"`
	UserId     interface{} //
	WalletAddr interface{} //
	ChainId    interface{} //
}
