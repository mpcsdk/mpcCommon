// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Chaincfg is the golang structure for table chaincfg.
type Chaincfg struct {
	Id         int         `json:"id"         ` // ID
	CreateTime *gtime.Time `json:"createTime" ` // 创建时间
	UpdateTime *gtime.Time `json:"updateTime" ` // 更新时间
	Coin       string      `json:"coin"       ` // 币种
	Rpc        string      `json:"rpc"        ` // RPC地址
	IsEnable   int         `json:"isEnable"   ` // 启用
	ChainId    int64       `json:"chainId"    ` // 链id
}
