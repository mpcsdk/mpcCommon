// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AggNftDao is the data access object for table aggNft.
type AggNftDao struct {
	table   string        // table is the underlying table name of the DAO.
	group   string        // group is the database configuration group name of current DAO.
	columns AggNftColumns // columns contains all the column names of Table for convenient usage.
}

// AggNftColumns defines and stores column names for table aggNft.
type AggNftColumns struct {
	TxHash     string //
	FromAddr   string //
	ToAddr     string //
	Contract   string //
	Value      string //
	StartTs    string //
	ChainId    string //
	EndTs      string //
	StartBlock string //
	EndBlock   string //
}

// aggNftColumns holds the columns for table aggNft.
var aggNftColumns = AggNftColumns{
	TxHash:     "tx_hash",
	FromAddr:   "from_addr",
	ToAddr:     "to_addr",
	Contract:   "contract",
	Value:      "value",
	StartTs:    "start_ts",
	ChainId:    "chain_id",
	EndTs:      "end_ts",
	StartBlock: "start_block",
	EndBlock:   "end_block",
}

// NewAggNftDao creates and returns a new DAO object for table data access.
func NewAggNftDao() *AggNftDao {
	return &AggNftDao{
		group:   "default",
		table:   "aggNft",
		columns: aggNftColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AggNftDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AggNftDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AggNftDao) Columns() AggNftColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AggNftDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AggNftDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AggNftDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
