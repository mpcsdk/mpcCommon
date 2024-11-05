// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// RelayerdminRelayerFeeDao is the data access object for table relayer_fee.
type RelayerdminRelayerFeeDao struct {
	table   string                       // table is the underlying table name of the DAO.
	group   string                       // group is the database configuration group name of current DAO.
	columns RelayerdminRelayerFeeColumns // columns contains all the column names of Table for convenient usage.
}

// RelayerdminRelayerFeeColumns defines and stores column names for table relayer_fee.
type RelayerdminRelayerFeeColumns struct {
	Id          string // ID
	CreateTime  string // 创建时间
	UpdateTime  string // 更新时间
	ChainId     string //
	ChannelId   string // 渠道号
	ChannelName string // 渠道名
	Fee         string // fee
	IsEnable    string // 启用
}

// relayerdminRelayerFeeColumns holds the columns for table relayer_fee.
var relayerdminRelayerFeeColumns = RelayerdminRelayerFeeColumns{
	Id:          "id",
	CreateTime:  "createTime",
	UpdateTime:  "updateTime",
	ChainId:     "chainId",
	ChannelId:   "channelId",
	ChannelName: "channelName",
	Fee:         "fee",
	IsEnable:    "isEnable",
}

// NewRelayerdminRelayerFeeDao creates and returns a new DAO object for table data access.
func NewRelayerdminRelayerFeeDao() *RelayerdminRelayerFeeDao {
	return &RelayerdminRelayerFeeDao{
		group:   "relayeradmin",
		table:   "relayer_fee",
		columns: relayerdminRelayerFeeColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *RelayerdminRelayerFeeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *RelayerdminRelayerFeeDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *RelayerdminRelayerFeeDao) Columns() RelayerdminRelayerFeeColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *RelayerdminRelayerFeeDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *RelayerdminRelayerFeeDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *RelayerdminRelayerFeeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
