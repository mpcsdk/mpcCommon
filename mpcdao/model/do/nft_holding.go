// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// NftHolding is the golang structure of table nft_holding for DAO operations like Where/Data.
type NftHolding struct {
	g.Meta      `orm:"table:nft_holding, do:true"`
	ChainId     interface{} //
	Address     interface{} //
	Contract    interface{} //
	TokenId     interface{} //
	Value       interface{} //
	BlockNumber interface{} //
	UpdatedAt   *gtime.Time //
	Kind        interface{} //
}
