// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// RiskadminChaincfgDao is the data access object for table chaincfg.
type RiskadminChaincfgDao struct {
	table   string                   // table is the underlying table name of the DAO.
	group   string                   // group is the database configuration group name of current DAO.
	columns RiskadminChaincfgColumns // columns contains all the column names of Table for convenient usage.
}

// RiskadminChaincfgColumns defines and stores column names for table chaincfg.
type RiskadminChaincfgColumns struct {
	Id         string // ID
	CreateTime string // 创建时间
	UpdateTime string // 更新时间
	Coin       string // 币种
	Rpc        string // RPC地址
	IsEnable   string // 启用
	ChainId    string // 链id
	Heigh      string // 高度
	Decimal    string // decimal
}

// riskadminChaincfgColumns holds the columns for table chaincfg.
var riskadminChaincfgColumns = RiskadminChaincfgColumns{
	Id:         "id",
	CreateTime: "createTime",
	UpdateTime: "updateTime",
	Coin:       "coin",
	Rpc:        "rpc",
	IsEnable:   "isEnable",
	ChainId:    "chainId",
	Heigh:      "heigh",
	Decimal:    "decimal",
}

// NewRiskadminChaincfgDao creates and returns a new DAO object for table data access.
func NewRiskadminChaincfgDao() *RiskadminChaincfgDao {
	return &RiskadminChaincfgDao{
		group:   "riskctrl",
		table:   "chaincfg",
		columns: riskadminChaincfgColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *RiskadminChaincfgDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *RiskadminChaincfgDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *RiskadminChaincfgDao) Columns() RiskadminChaincfgColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *RiskadminChaincfgDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *RiskadminChaincfgDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *RiskadminChaincfgDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
