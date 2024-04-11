// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ChaincfgDao is the data access object for table chaincfg.
type ChaincfgDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns ChaincfgColumns // columns contains all the column names of Table for convenient usage.
}

// ChaincfgColumns defines and stores column names for table chaincfg.
type ChaincfgColumns struct {
	Id         string // ID
	CreateTime string // 创建时间
	UpdateTime string // 更新时间
	Coin       string // 币种
	Rpc        string // RPC地址
	IsEnable   string // 启用
	ChainId    string // 链id
}

// chaincfgColumns holds the columns for table chaincfg.
var chaincfgColumns = ChaincfgColumns{
	Id:         "id",
	CreateTime: "createTime",
	UpdateTime: "updateTime",
	Coin:       "coin",
	Rpc:        "rpc",
	IsEnable:   "isEnable",
	ChainId:    "chainId",
}

// NewChaincfgDao creates and returns a new DAO object for table data access.
func NewChaincfgDao() *ChaincfgDao {
	return &ChaincfgDao{
		group:   "riskcontrol",
		table:   "chaincfg",
		columns: chaincfgColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ChaincfgDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ChaincfgDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ChaincfgDao) Columns() ChaincfgColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ChaincfgDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ChaincfgDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ChaincfgDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
