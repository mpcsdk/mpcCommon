// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// RelayeradminSpecifiedGas is the golang structure of table specifiedGas for DAO operations like Where/Data.
type RelayeradminSpecifiedGas struct {
	g.Meta       `orm:"table:specifiedGas, do:true"`
	Id           interface{} // ID
	CreateTime   *gtime.Time // 创建时间
	UpdateTime   *gtime.Time // 更新时间
	ContractAddr interface{} // 合约地址
	MethodSig    interface{} // 方法签名
	IsEnable     interface{} // 启用
	GasUsed      interface{} // gas
	ChainId      interface{} // gas
	Desc         interface{} // 描述
}
