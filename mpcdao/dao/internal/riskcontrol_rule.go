// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// RiskcontrolRuleDao is the data access object for table riskcontrol_rule.
type RiskcontrolRuleDao struct {
	table   string                 // table is the underlying table name of the DAO.
	group   string                 // group is the database configuration group name of current DAO.
	columns RiskcontrolRuleColumns // columns contains all the column names of Table for convenient usage.
}

// RiskcontrolRuleColumns defines and stores column names for table riskcontrol_rule.
type RiskcontrolRuleColumns struct {
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

// riskcontrolRuleColumns holds the columns for table riskcontrol_rule.
var riskcontrolRuleColumns = RiskcontrolRuleColumns{
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

// NewRiskcontrolRuleDao creates and returns a new DAO object for table data access.
func NewRiskcontrolRuleDao() *RiskcontrolRuleDao {
	return &RiskcontrolRuleDao{
		group:   "riskcontrol",
		table:   "riskcontrol_rule",
		columns: riskcontrolRuleColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *RiskcontrolRuleDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *RiskcontrolRuleDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *RiskcontrolRuleDao) Columns() RiskcontrolRuleColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *RiskcontrolRuleDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *RiskcontrolRuleDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *RiskcontrolRuleDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
