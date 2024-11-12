// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// RelayerRelayerChannelFeeDao is the data access object for table relayer_channel_fee.
type RelayerRelayerChannelFeeDao struct {
	table   string                          // table is the underlying table name of the DAO.
	group   string                          // group is the database configuration group name of current DAO.
	columns RelayerRelayerChannelFeeColumns // columns contains all the column names of Table for convenient usage.
}

// RelayerRelayerChannelFeeColumns defines and stores column names for table relayer_channel_fee.
type RelayerRelayerChannelFeeColumns struct {
}

// relayerRelayerChannelFeeColumns holds the columns for table relayer_channel_fee.
var relayerRelayerChannelFeeColumns = RelayerRelayerChannelFeeColumns{}

// NewRelayerRelayerChannelFeeDao creates and returns a new DAO object for table data access.
func NewRelayerRelayerChannelFeeDao() *RelayerRelayerChannelFeeDao {
	return &RelayerRelayerChannelFeeDao{
		group:   "relayer",
		table:   "relayer_channel_fee",
		columns: relayerRelayerChannelFeeColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *RelayerRelayerChannelFeeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *RelayerRelayerChannelFeeDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *RelayerRelayerChannelFeeDao) Columns() RelayerRelayerChannelFeeColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *RelayerRelayerChannelFeeDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *RelayerRelayerChannelFeeDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *RelayerRelayerChannelFeeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
