// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Contractabi is the golang structure for table contractabi.
type Contractabi struct {
	Id              int         `json:"id"              orm:"id"              ` // ID
	CreateTime      *gtime.Time `json:"createTime"      orm:"createTime"      ` // 创建时间
	UpdateTime      *gtime.Time `json:"updateTime"      orm:"updateTime"      ` // 更新时间
	ContractName    string      `json:"contractName"    orm:"contractName"    ` // 合约名
	ContractAddress string      `json:"contractAddress" orm:"contractAddress" ` // 合约地址
	SceneNo         string      `json:"sceneNo"         orm:"sceneNo"         ` // 场景号
	AbiContent      string      `json:"abiContent"      orm:"abiContent"      ` // 合约abi
	ContractKind    string      `json:"contractKind"    orm:"contractKind"    ` // 合约类型
	ChainId         int64       `json:"chainId"         orm:"chainId"         ` // 链id
	Decimal         int         `json:"decimal"         orm:"decimal"         ` //
	TokenId         string      `json:"tokenId"         orm:"tokenId"         ` // tokenId
}
