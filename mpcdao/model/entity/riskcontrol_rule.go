// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// RiskcontrolRule is the golang structure for table riskcontrol_rule.
type RiskcontrolRule struct {
	Id         int         `json:"id"         ` // ID
	CreateTime *gtime.Time `json:"createTime" ` // 创建时间
	UpdateTime *gtime.Time `json:"updateTime" ` // 更新时间
	SceneNo    string      `json:"sceneNo"    ` // 场景号
	RuleName   string      `json:"ruleName"   ` // 风控名
	RuleStr    string      `json:"ruleStr"    ` // 脚本
	Salience   int         `json:"salience"   ` // 优先级
	IsEnable   int         `json:"isEnable"   ` // 启用
	Desc       string      `json:"desc"       ` // 描述
	ChainId    string      `json:"chainId"    ` // 链id
}
