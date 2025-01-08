// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// RiskadminChaincfg is the golang structure for table riskadmin_chaincfg.
type RiskadminChaincfg struct {
	Id         int         `json:"id"         orm:"id"         ` // ID
	CreateTime *gtime.Time `json:"createTime" orm:"createTime" ` // 创建时间
	UpdateTime *gtime.Time `json:"updateTime" orm:"updateTime" ` // 更新时间
	Coin       string      `json:"coin"       orm:"coin"       ` // 币种
	Rpc        string      `json:"rpc"        orm:"rpc"        ` // RPC地址
	IsEnable   int         `json:"isEnable"   orm:"isEnable"   ` // 启用
	ChainId    int64       `json:"chainId"    orm:"chainId"    ` // 链id
	Heigh      int64       `json:"heigh"      orm:"heigh"      ` // 高度
}
