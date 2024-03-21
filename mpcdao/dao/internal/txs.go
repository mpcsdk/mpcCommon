// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TxsDao is the data access object for table txs.
type TxsDao struct {
	table   string     // table is the underlying table name of the DAO.
	group   string     // group is the database configuration group name of current DAO.
	columns TxsColumns // columns contains all the column names of Table for convenient usage.
}

// TxsColumns defines and stores column names for table txs.
type TxsColumns struct {
	ChainId     string //
	TxHash      string //
	TxIndex     string //
	BlockHash   string //
	BlockNumber string //
	From        string //
	To          string //
	Value       string //
	Contract    string //
	Data        string //
	LogIndex    string //
}

// txsColumns holds the columns for table txs.
var txsColumns = TxsColumns{
	ChainId:     "chain_id",
	TxHash:      "tx_hash",
	TxIndex:     "tx_index",
	BlockHash:   "block_hash",
	BlockNumber: "block_number",
	From:        "from",
	To:          "to",
	Value:       "value",
	Contract:    "contract",
	Data:        "data",
	LogIndex:    "log_index",
}

// NewTxsDao creates and returns a new DAO object for table data access.
func NewTxsDao() *TxsDao {
	return &TxsDao{
		group:   "default",
		table:   "txs",
		columns: txsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TxsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TxsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TxsDao) Columns() TxsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TxsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TxsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TxsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
