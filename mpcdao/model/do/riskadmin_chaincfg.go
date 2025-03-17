// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// RiskadminChaincfg is the golang structure of table chaincfg for DAO operations like Where/Data.
type RiskadminChaincfg struct {
	g.Meta     `orm:"table:chaincfg, do:true"`
	Id         interface{} // ID
	CreateTime *gtime.Time // 创建时间
	UpdateTime *gtime.Time // 更新时间
	Coin       interface{} // 币种
	Rpc        interface{} // RPC地址
	IsEnable   interface{} // 启用
	ChainId    interface{} // 链id
	Heigh      interface{} // 高度
	Decimal    interface{} // decimal
}
