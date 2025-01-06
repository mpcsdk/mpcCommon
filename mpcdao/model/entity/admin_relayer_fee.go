// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminRelayerFee is the golang structure for table admin_relayer_fee.
type AdminRelayerFee struct {
	Id         int         `json:"id"         orm:"id"         ` // ID
	CreateTime *gtime.Time `json:"createTime" orm:"createTime" ` // 创建时间
	UpdateTime *gtime.Time `json:"updateTime" orm:"updateTime" ` // 更新时间
	ChainId    int         `json:"chainId"    orm:"chainId"    ` //
	AppId      string      `json:"appId"      orm:"appId"      ` // 渠道号
	AppName    string      `json:"appName"    orm:"appName"    ` // 渠道名
	Fee        float64     `json:"fee"        orm:"fee"        ` // fee
	IsEnable   int         `json:"isEnable"   orm:"isEnable"   ` // 启用
}
