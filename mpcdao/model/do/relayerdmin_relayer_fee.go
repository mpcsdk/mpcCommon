// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// RelayerdminRelayerFee is the golang structure of table relayer_fee for DAO operations like Where/Data.
type RelayerdminRelayerFee struct {
	g.Meta      `orm:"table:relayer_fee, do:true"`
	Id          interface{} // ID
	CreateTime  *gtime.Time // 创建时间
	UpdateTime  *gtime.Time // 更新时间
	ChainId     interface{} //
	ChannelId   interface{} // 渠道号
	ChannelName interface{} // 渠道名
	Fee         interface{} // fee
	IsEnable    interface{} // 启用
}
