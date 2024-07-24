// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Contractrule is the golang structure for table contractrule.
type Contractrule struct {
	Id               int         `json:"id"               orm:"id"               ` // ID
	CreateTime       *gtime.Time `json:"createTime"       orm:"createTime"       ` // 创建时间
	UpdateTime       *gtime.Time `json:"updateTime"       orm:"updateTime"       ` // 更新时间
	SceneNo          string      `json:"sceneNo"          orm:"sceneNo"          ` // 场景号
	ContractAddress  string      `json:"contractAddress"  orm:"contractAddress"  ` // 合约地址
	ContractName     string      `json:"contractName"     orm:"contractName"     ` // 合约名
	MethodName       string      `json:"methodName"       orm:"methodName"       ` // 方法名
	MethodSignature  string      `json:"methodSignature"  orm:"methodSignature"  ` // 方法签名
	MethodFromField  string      `json:"methodFromField"  orm:"methodFromField"  ` // 方法from字段名
	MethodToField    string      `json:"methodToField"    orm:"methodToField"    ` // 方法to字段名
	MethodValueField string      `json:"methodValueField" orm:"methodValueField" ` // 方法value字段名
	ContractKind     string      `json:"contractKind"     orm:"contractKind"     ` // 合约类型
	WhiteAddrList    string      `json:"whiteAddrList"    orm:"whiteAddrList"    ` // to地址白名单
	ChainId          int64       `json:"chainId"          orm:"chainId"          ` // 链id
}
