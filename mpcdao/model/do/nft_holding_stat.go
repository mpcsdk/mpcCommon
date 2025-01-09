// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// NftHoldingStat is the golang structure of table nft_holding_stat for DAO operations like Where/Data.
type NftHoldingStat struct {
	g.Meta      `orm:"table:nft_holding_stat, do:true"`
	ChainId     interface{} //
	BlockNumber interface{} //
	UpdatedAt   *gtime.Time //
}
