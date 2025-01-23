// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// RelayeradminAssignFee is the golang structure for table relayeradmin_assignFee.
type RelayeradminAssignFee struct {
	Id         int         `json:"id"         orm:"id"         ` // ID
	CreateTime *gtime.Time `json:"createTime" orm:"createTime" ` // 创建时间
	UpdateTime *gtime.Time `json:"updateTime" orm:"updateTime" ` // 更新时间
	AppId      string      `json:"appId"      orm:"appId"      ` // 渠道号
	AppName    string      `json:"appName"    orm:"appName"    ` // 渠道名
	Fee        float64     `json:"fee"        orm:"fee"        ` // fee
	IsEnable   int         `json:"isEnable"   orm:"isEnable"   ` // 启用
	ChainId    int64       `json:"chainId"    orm:"chainId"    ` // gas
}
