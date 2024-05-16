// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// FcmOfflineMsgDao is the data access object for table fcm_offline_msg.
type FcmOfflineMsgDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of current DAO.
	columns FcmOfflineMsgColumns // columns contains all the column names of Table for convenient usage.
}

// FcmOfflineMsgColumns defines and stores column names for table fcm_offline_msg.
type FcmOfflineMsgColumns struct {
	FmcToken    string //
	Title       string //
	Body        string //
	Data        string //
	Err         string //
	Address     string //
	UserId      string //
	CreatedTime string //
	Id          string //
}

// fcmOfflineMsgColumns holds the columns for table fcm_offline_msg.
var fcmOfflineMsgColumns = FcmOfflineMsgColumns{
	FmcToken:    "fmc_token",
	Title:       "title",
	Body:        "body",
	Data:        "data",
	Err:         "err",
	Address:     "address",
	UserId:      "user_id",
	CreatedTime: "created_time",
	Id:          "id",
}

// NewFcmOfflineMsgDao creates and returns a new DAO object for table data access.
func NewFcmOfflineMsgDao() *FcmOfflineMsgDao {
	return &FcmOfflineMsgDao{
		group:   "fcm",
		table:   "fcm_offline_msg",
		columns: fcmOfflineMsgColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *FcmOfflineMsgDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *FcmOfflineMsgDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *FcmOfflineMsgDao) Columns() FcmOfflineMsgColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *FcmOfflineMsgDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *FcmOfflineMsgDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *FcmOfflineMsgDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
