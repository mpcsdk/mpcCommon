// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ChainTransferDao is the data access object for table chain_transfer.
type ChainTransferDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of current DAO.
	columns ChainTransferColumns // columns contains all the column names of Table for convenient usage.
}

// ChainTransferColumns defines and stores column names for table chain_transfer.
type ChainTransferColumns struct {
	ChainId   string //
	Height    string //
	BlockHash string //
	Ts        string //
	TxHash    string //
	TxIdx     string //
	LogIdx    string //
	From      string //
	To        string //
	Contract  string //
	Value     string //
	Gas       string //
	GasPrice  string //
	Nonce     string //
	Kind      string //
	TokenId   string //
	Removed   string //
	Status    string //
}

// chainTransferColumns holds the columns for table chain_transfer.
var chainTransferColumns = ChainTransferColumns{
	ChainId:   "chain_id",
	Height:    "height",
	BlockHash: "block_hash",
	Ts:        "ts",
	TxHash:    "tx_hash",
	TxIdx:     "tx_idx",
	LogIdx:    "log_idx",
	From:      "from",
	To:        "to",
	Contract:  "contract",
	Value:     "value",
	Gas:       "gas",
	GasPrice:  "gas_price",
	Nonce:     "nonce",
	Kind:      "kind",
	TokenId:   "token_id",
	Removed:   "removed",
	Status:    "status",
}

// NewChainTransferDao creates and returns a new DAO object for table data access.
func NewChainTransferDao() *ChainTransferDao {
	return &ChainTransferDao{
		group:   "sync_chain",
		table:   "chain_transfer",
		columns: chainTransferColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ChainTransferDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ChainTransferDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ChainTransferDao) Columns() ChainTransferColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ChainTransferDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ChainTransferDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ChainTransferDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
