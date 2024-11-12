// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// RelayerReceiptDao is the data access object for table receipt.
type RelayerReceiptDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of current DAO.
	columns RelayerReceiptColumns // columns contains all the column names of Table for convenient usage.
}

// RelayerReceiptColumns defines and stores column names for table receipt.
type RelayerReceiptColumns struct {
	RelayerIdent      string //
	RelayerTxHash     string //
	ChainId           string //
	BlockHash         string //
	BlockNumber       string //
	TxIdx             string //
	TxHash            string //
	CumulativeGasUsed string //
	GasUsed           string //
	EffectiveGasPrice string //
	ContractAddress   string //
	Type              string //
	Root              string //
	Bloom             string //
	Logs              string //
	Status            string //
	TransactionFee    string //
	Createdat         string //
	AppId             string //
	UserId            string //
	Payer             string //
}

// relayerReceiptColumns holds the columns for table receipt.
var relayerReceiptColumns = RelayerReceiptColumns{
	RelayerIdent:      "relayer_ident",
	RelayerTxHash:     "relayer_tx_hash",
	ChainId:           "chain_id",
	BlockHash:         "block_hash",
	BlockNumber:       "block_number",
	TxIdx:             "tx_idx",
	TxHash:            "tx_hash",
	CumulativeGasUsed: "cumulative_gas_used",
	GasUsed:           "gas_used",
	EffectiveGasPrice: "effective_gas_price",
	ContractAddress:   "contract_address",
	Type:              "type",
	Root:              "root",
	Bloom:             "bloom",
	Logs:              "logs",
	Status:            "status",
	TransactionFee:    "transaction_fee",
	Createdat:         "createdat",
	AppId:             "app_id",
	UserId:            "user_id",
	Payer:             "payer",
}

// NewRelayerReceiptDao creates and returns a new DAO object for table data access.
func NewRelayerReceiptDao() *RelayerReceiptDao {
	return &RelayerReceiptDao{
		group:   "relayer",
		table:   "receipt",
		columns: relayerReceiptColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *RelayerReceiptDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *RelayerReceiptDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *RelayerReceiptDao) Columns() RelayerReceiptColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *RelayerReceiptDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *RelayerReceiptDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *RelayerReceiptDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
