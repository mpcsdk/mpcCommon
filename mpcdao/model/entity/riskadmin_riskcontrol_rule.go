// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// RiskadminRiskcontrolRule is the golang structure for table riskadmin_riskcontrol_rule.
type RiskadminRiskcontrolRule struct {
	Id         int         `json:"id"         orm:"id"         ` // ID
	CreateTime *gtime.Time `json:"createTime" orm:"createTime" ` // 创建时间
	UpdateTime *gtime.Time `json:"updateTime" orm:"updateTime" ` // 更新时间
	SceneNo    string      `json:"sceneNo"    orm:"sceneNo"    ` // 场景号
	RuleName   string      `json:"ruleName"   orm:"ruleName"   ` // 风控名
	RuleStr    string      `json:"ruleStr"    orm:"ruleStr"    ` // 脚本
	Salience   int         `json:"salience"   orm:"salience"   ` // 优先级
	Desc       string      `json:"desc"       orm:"desc"       ` // 描述
	IsEnable   int         `json:"isEnable"   orm:"isEnable"   ` // 启用
}
