// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// RiskadminRiskcontrolRuleDao is the data access object for table riskcontrol_rule.
type RiskadminRiskcontrolRuleDao struct {
	table   string                          // table is the underlying table name of the DAO.
	group   string                          // group is the database configuration group name of current DAO.
	columns RiskadminRiskcontrolRuleColumns // columns contains all the column names of Table for convenient usage.
}

// RiskadminRiskcontrolRuleColumns defines and stores column names for table riskcontrol_rule.
type RiskadminRiskcontrolRuleColumns struct {
	Id         string // ID
	CreateTime string // 创建时间
	UpdateTime string // 更新时间
	SceneNo    string // 场景号
	RuleName   string // 风控名
	RuleStr    string // 脚本
	Salience   string // 优先级
	Desc       string // 描述
	IsEnable   string // 启用
}

// riskadminRiskcontrolRuleColumns holds the columns for table riskcontrol_rule.
var riskadminRiskcontrolRuleColumns = RiskadminRiskcontrolRuleColumns{
	Id:         "id",
	CreateTime: "createTime",
	UpdateTime: "updateTime",
	SceneNo:    "sceneNo",
	RuleName:   "ruleName",
	RuleStr:    "ruleStr",
	Salience:   "salience",
	Desc:       "desc",
	IsEnable:   "isEnable",
}

// NewRiskadminRiskcontrolRuleDao creates and returns a new DAO object for table data access.
func NewRiskadminRiskcontrolRuleDao() *RiskadminRiskcontrolRuleDao {
	return &RiskadminRiskcontrolRuleDao{
		group:   "riskctrl",
		table:   "riskcontrol_rule",
		columns: riskadminRiskcontrolRuleColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *RiskadminRiskcontrolRuleDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *RiskadminRiskcontrolRuleDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *RiskadminRiskcontrolRuleDao) Columns() RiskadminRiskcontrolRuleColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *RiskadminRiskcontrolRuleDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *RiskadminRiskcontrolRuleDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *RiskadminRiskcontrolRuleDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
