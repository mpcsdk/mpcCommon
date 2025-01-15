// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// RelayeradminAssignFeeDao is the data access object for table assignFee.
type RelayeradminAssignFeeDao struct {
	table   string                       // table is the underlying table name of the DAO.
	group   string                       // group is the database configuration group name of current DAO.
	columns RelayeradminAssignFeeColumns // columns contains all the column names of Table for convenient usage.
}

// RelayeradminAssignFeeColumns defines and stores column names for table assignFee.
type RelayeradminAssignFeeColumns struct {
	Id         string // ID
	CreateTime string // 创建时间
	UpdateTime string // 更新时间
	ChainId    string //
	AppId      string // 渠道号
	AppName    string // 渠道名
	Fee        string // fee
	IsEnable   string // 启用
}

// relayeradminAssignFeeColumns holds the columns for table assignFee.
var relayeradminAssignFeeColumns = RelayeradminAssignFeeColumns{
	Id:         "id",
	CreateTime: "createTime",
	UpdateTime: "updateTime",
	ChainId:    "chainId",
	AppId:      "appId",
	AppName:    "appName",
	Fee:        "fee",
	IsEnable:   "isEnable",
}

// NewRelayeradminAssignFeeDao creates and returns a new DAO object for table data access.
func NewRelayeradminAssignFeeDao() *RelayeradminAssignFeeDao {
	return &RelayeradminAssignFeeDao{
		group:   "relayeradmin",
		table:   "assignFee",
		columns: relayeradminAssignFeeColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *RelayeradminAssignFeeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *RelayeradminAssignFeeDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *RelayeradminAssignFeeDao) Columns() RelayeradminAssignFeeColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *RelayeradminAssignFeeDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *RelayeradminAssignFeeDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *RelayeradminAssignFeeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
