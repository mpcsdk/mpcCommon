// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// RelayerChannelFee is the golang structure of table channel_fee for DAO operations like Where/Data.
type RelayerChannelFee struct {
	g.Meta    `orm:"table:channel_fee, do:true"`
	AppId     interface{} //
	Fee       interface{} //
	Updatedat *gtime.Time //
}
