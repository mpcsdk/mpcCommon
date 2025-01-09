// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// RiskcontrolRule is the golang structure of table riskcontrol_rule for DAO operations like Where/Data.
type RiskcontrolRule struct {
	g.Meta     `orm:"table:riskcontrol_rule, do:true"`
	Id         interface{} // ID
	CreateTime *gtime.Time // 创建时间
	UpdateTime *gtime.Time // 更新时间
	SceneNo    interface{} // 场景号
	RuleName   interface{} // 风控名
	RuleStr    interface{} // 脚本
	Salience   interface{} // 优先级
	Desc       interface{} // 描述
	IsEnable   interface{} // 启用
}
