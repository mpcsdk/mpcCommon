// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AdminRelayerAppDao is the data access object for table relayer_app.
type AdminRelayerAppDao struct {
	table   string                 // table is the underlying table name of the DAO.
	group   string                 // group is the database configuration group name of current DAO.
	columns AdminRelayerAppColumns // columns contains all the column names of Table for convenient usage.
}

// AdminRelayerAppColumns defines and stores column names for table relayer_app.
type AdminRelayerAppColumns struct {
	Id         string // ID
	CreateTime string // 创建时间
	UpdateTime string // 更新时间
	AppId      string // 渠道号
	AppName    string // 渠道名
	IsEnable   string // 启用
}

// adminRelayerAppColumns holds the columns for table relayer_app.
var adminRelayerAppColumns = AdminRelayerAppColumns{
	Id:         "id",
	CreateTime: "createTime",
	UpdateTime: "updateTime",
	AppId:      "appId",
	AppName:    "appName",
	IsEnable:   "isEnable",
}

// NewAdminRelayerAppDao creates and returns a new DAO object for table data access.
func NewAdminRelayerAppDao() *AdminRelayerAppDao {
	return &AdminRelayerAppDao{
		group:   "relayeradmin",
		table:   "relayer_app",
		columns: adminRelayerAppColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AdminRelayerAppDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AdminRelayerAppDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AdminRelayerAppDao) Columns() AdminRelayerAppColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AdminRelayerAppDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AdminRelayerAppDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AdminRelayerAppDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
