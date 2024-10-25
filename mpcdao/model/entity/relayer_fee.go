// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// RelayerFee is the golang structure for table relayer_fee.
type RelayerFee struct {
	Id          int         `json:"id"          orm:"id"          ` // ID
	CreateTime  *gtime.Time `json:"createTime"  orm:"createTime"  ` // 创建时间
	UpdateTime  *gtime.Time `json:"updateTime"  orm:"updateTime"  ` // 更新时间
	ChainId     int         `json:"chainId"     orm:"chainId"     ` //
	ChannelId   string      `json:"channelId"   orm:"channelId"   ` // 渠道号
	ChannelName string      `json:"channelName" orm:"channelName" ` // 渠道名
	Fee         float64     `json:"fee"         orm:"fee"         ` // fee
}
