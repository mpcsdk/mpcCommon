// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// RelayerChannel is the golang structure for table relayer_channel.
type RelayerChannel struct {
	Id          int         `json:"id"          orm:"id"          ` // ID
	CreateTime  *gtime.Time `json:"createTime"  orm:"createTime"  ` // 创建时间
	UpdateTime  *gtime.Time `json:"updateTime"  orm:"updateTime"  ` // 更新时间
	ChannelId   string      `json:"channelId"   orm:"channelId"   ` // 渠道号
	ChannelName string      `json:"channelName" orm:"channelName" ` // 渠道名
	IsEnable    int         `json:"isEnable"    orm:"isEnable"    ` // 启用
}
