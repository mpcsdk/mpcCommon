// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AggFtDao is the data access object for table aggFt.
type AggFtDao struct {
	table   string       // table is the underlying table name of the DAO.
	group   string       // group is the database configuration group name of current DAO.
	columns AggFtColumns // columns contains all the column names of Table for convenient usage.
}

// AggFtColumns defines and stores column names for table aggFt.
type AggFtColumns struct {
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

// aggFtColumns holds the columns for table aggFt.
var aggFtColumns = AggFtColumns{
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

// NewAggFtDao creates and returns a new DAO object for table data access.
func NewAggFtDao() *AggFtDao {
	return &AggFtDao{
		group:   "default",
		table:   "aggFt",
		columns: aggFtColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AggFtDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AggFtDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AggFtDao) Columns() AggFtColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AggFtDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AggFtDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AggFtDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
