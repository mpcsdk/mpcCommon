// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// RelayeradminSpecifiedGas is the golang structure for table relayeradmin_specifiedGas.
type RelayeradminSpecifiedGas struct {
	Id           int         `json:"id"           orm:"id"           ` // ID
	CreateTime   *gtime.Time `json:"createTime"   orm:"createTime"   ` // 创建时间
	UpdateTime   *gtime.Time `json:"updateTime"   orm:"updateTime"   ` // 更新时间
	ContractAddr string      `json:"contractAddr" orm:"contractAddr" ` // 合约地址
	MethodSig    string      `json:"methodSig"    orm:"methodSig"    ` // 方法签名
	IsEnable     int         `json:"isEnable"     orm:"isEnable"     ` // 启用
	GasUsed      int64       `json:"gasUsed"      orm:"gasUsed"      ` // gas
	ChainId      int64       `json:"chainId"      orm:"chainId"      ` // gas
	Desc         string      `json:"desc"         orm:"desc"         ` // 描述
}
