// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// WalletAddrDao is the data access object for table wallet_addr.
type WalletAddrDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns WalletAddrColumns // columns contains all the column names of Table for convenient usage.
}

// WalletAddrColumns defines and stores column names for table wallet_addr.
type WalletAddrColumns struct {
	UserId     string //
	WalletAddr string //
	ChainId    string //
}

// walletAddrColumns holds the columns for table wallet_addr.
var walletAddrColumns = WalletAddrColumns{
	UserId:     "user_id",
	WalletAddr: "wallet_addr",
	ChainId:    "chain_id",
}

// NewWalletAddrDao creates and returns a new DAO object for table data access.
func NewWalletAddrDao() *WalletAddrDao {
	return &WalletAddrDao{
		group:   "mpc",
		table:   "wallet_addr",
		columns: walletAddrColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *WalletAddrDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *WalletAddrDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *WalletAddrDao) Columns() WalletAddrColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *WalletAddrDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *WalletAddrDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *WalletAddrDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
