// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// NftHoldingStatDao is the data access object for table nft_holding_stat.
type NftHoldingStatDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of current DAO.
	columns NftHoldingStatColumns // columns contains all the column names of Table for convenient usage.
}

// NftHoldingStatColumns defines and stores column names for table nft_holding_stat.
type NftHoldingStatColumns struct {
	ChainId     string //
	BlockNumber string //
	UpdatedAt   string //
}

// nftHoldingStatColumns holds the columns for table nft_holding_stat.
var nftHoldingStatColumns = NftHoldingStatColumns{
	ChainId:     "chain_id",
	BlockNumber: "block_number",
	UpdatedAt:   "updated_at",
}

// NewNftHoldingStatDao creates and returns a new DAO object for table data access.
func NewNftHoldingStatDao() *NftHoldingStatDao {
	return &NftHoldingStatDao{
		group:   "sync_nft_holding",
		table:   "nft_holding_stat",
		columns: nftHoldingStatColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *NftHoldingStatDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *NftHoldingStatDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *NftHoldingStatDao) Columns() NftHoldingStatColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *NftHoldingStatDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *NftHoldingStatDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *NftHoldingStatDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
