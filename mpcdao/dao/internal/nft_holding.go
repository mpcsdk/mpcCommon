// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// NftHoldingDao is the data access object for table nft_holding.
type NftHoldingDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns NftHoldingColumns // columns contains all the column names of Table for convenient usage.
}

// NftHoldingColumns defines and stores column names for table nft_holding.
type NftHoldingColumns struct {
	ChainId     string //
	Address     string //
	Contract    string //
	TokenId     string //
	Value       string //
	BlockNumber string //
	UpdatedAt   string //
	Kind        string //
}

// nftHoldingColumns holds the columns for table nft_holding.
var nftHoldingColumns = NftHoldingColumns{
	ChainId:     "chain_id",
	Address:     "address",
	Contract:    "contract",
	TokenId:     "token_id",
	Value:       "value",
	BlockNumber: "block_number",
	UpdatedAt:   "updated_at",
	Kind:        "kind",
}

// NewNftHoldingDao creates and returns a new DAO object for table data access.
func NewNftHoldingDao() *NftHoldingDao {
	return &NftHoldingDao{
		group:   "sync_nft_holding",
		table:   "nft_holding",
		columns: nftHoldingColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *NftHoldingDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *NftHoldingDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *NftHoldingDao) Columns() NftHoldingColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *NftHoldingDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *NftHoldingDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *NftHoldingDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
