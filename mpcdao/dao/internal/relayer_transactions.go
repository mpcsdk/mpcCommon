// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// RelayerTransactionsDao is the data access object for table transactions.
type RelayerTransactionsDao struct {
	table   string                     // table is the underlying table name of the DAO.
	group   string                     // group is the database configuration group name of current DAO.
	columns RelayerTransactionsColumns // columns contains all the column names of Table for convenient usage.
}

// RelayerTransactionsColumns defines and stores column names for table transactions.
type RelayerTransactionsColumns struct {
	RelayerIdent    string //
	ChainId         string // 链id
	AppId           string //
	BlockNumber     string //
	RelayerTxHash   string // relayer根据用户交易链下计算的交易hash
	TxHash          string // 链上交易的交易hash
	TransactionData string // input transaction
	Target          string //
	WalletAddress   string // 用户钱包地址
	WalletNonce     string //
	Submitter       string // address of submitter
	SubmitterNonce  string // the nonce of submitter
	GasLimit        string // 上链时计算的gas limit
	GasPrice        string // 上链时的gas price
	Discount        string // 折扣百分比，值0-100
	Status          string // 交易状态  @values 0 收到交易 ｜ 1 交易成功 ｜ 2 交易失败 | 3 交易丢弃
	ErrorReason     string // 交易失败
	Createdat       string //
	Updatedat       string //
	RelayerFee      string // relayerfee
}

// relayerTransactionsColumns holds the columns for table transactions.
var relayerTransactionsColumns = RelayerTransactionsColumns{
	RelayerIdent:    "relayer_ident",
	ChainId:         "chain_id",
	AppId:           "app_id",
	BlockNumber:     "block_number",
	RelayerTxHash:   "relayer_tx_hash",
	TxHash:          "tx_hash",
	TransactionData: "transaction_data",
	Target:          "target",
	WalletAddress:   "wallet_address",
	WalletNonce:     "wallet_nonce",
	Submitter:       "submitter",
	SubmitterNonce:  "submitter_nonce",
	GasLimit:        "gas_limit",
	GasPrice:        "gas_price",
	Discount:        "discount",
	Status:          "status",
	ErrorReason:     "error_reason",
	Createdat:       "createdat",
	Updatedat:       "updatedat",
	RelayerFee:      "relayer_fee",
}

// NewRelayerTransactionsDao creates and returns a new DAO object for table data access.
func NewRelayerTransactionsDao() *RelayerTransactionsDao {
	return &RelayerTransactionsDao{
		group:   "relayer",
		table:   "transactions",
		columns: relayerTransactionsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *RelayerTransactionsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *RelayerTransactionsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *RelayerTransactionsDao) Columns() RelayerTransactionsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *RelayerTransactionsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *RelayerTransactionsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *RelayerTransactionsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
