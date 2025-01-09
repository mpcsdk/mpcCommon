// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// FcmTokenDao is the data access object for table fcm_token.
type FcmTokenDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns FcmTokenColumns // columns contains all the column names of Table for convenient usage.
}

// FcmTokenColumns defines and stores column names for table fcm_token.
type FcmTokenColumns struct {
	UserId      string // 用户id
	FcmToken    string //
	Token       string //
	Address     string //
	CreatedTime string //
	UpdatedTime string //
}

// fcmTokenColumns holds the columns for table fcm_token.
var fcmTokenColumns = FcmTokenColumns{
	UserId:      "user_id",
	FcmToken:    "fcm_token",
	Token:       "token",
	Address:     "address",
	CreatedTime: "created_time",
	UpdatedTime: "updated_time",
}

// NewFcmTokenDao creates and returns a new DAO object for table data access.
func NewFcmTokenDao() *FcmTokenDao {
	return &FcmTokenDao{
		group:   "fcm",
		table:   "fcm_token",
		columns: fcmTokenColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *FcmTokenDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *FcmTokenDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *FcmTokenDao) Columns() FcmTokenColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *FcmTokenDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *FcmTokenDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *FcmTokenDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
