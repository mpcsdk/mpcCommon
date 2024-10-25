// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// RelayerChannel is the golang structure of table relayer_channel for DAO operations like Where/Data.
type RelayerChannel struct {
	g.Meta      `orm:"table:relayer_channel, do:true"`
	Id          interface{} // ID
	CreateTime  *gtime.Time // 创建时间
	UpdateTime  *gtime.Time // 更新时间
	ChannelId   interface{} // 渠道号
	ChannelName interface{} // 渠道名
	IsEnable    interface{} // 启用
}
