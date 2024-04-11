// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Contractrule is the golang structure of table contractrule for DAO operations like Where/Data.
type Contractrule struct {
	g.Meta           `orm:"table:contractrule, do:true"`
	Id               interface{} // ID
	CreateTime       *gtime.Time // 创建时间
	UpdateTime       *gtime.Time // 更新时间
	SceneNo          interface{} // 场景号
	ContractAddress  interface{} // 合约地址
	ContractName     interface{} // 合约名
	MethodName       interface{} // 方法名
	MethodSignature  interface{} // 方法签名
	MethodFromField  interface{} // 方法from字段名
	MethodToField    interface{} // 方法to字段名
	MethodValueField interface{} // 方法value字段名
	ContractKind     interface{} // 合约类型
	WhiteAddrList    interface{} // to地址白名单
	ChainId          interface{} // 链id
}
