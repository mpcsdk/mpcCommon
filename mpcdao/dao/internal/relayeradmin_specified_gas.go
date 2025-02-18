// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// RelayeradminSpecifiedGasDao is the data access object for table specifiedGas.
type RelayeradminSpecifiedGasDao struct {
	table   string                          // table is the underlying table name of the DAO.
	group   string                          // group is the database configuration group name of current DAO.
	columns RelayeradminSpecifiedGasColumns // columns contains all the column names of Table for convenient usage.
}

// RelayeradminSpecifiedGasColumns defines and stores column names for table specifiedGas.
type RelayeradminSpecifiedGasColumns struct {
	Id           string // ID
	CreateTime   string // 创建时间
	UpdateTime   string // 更新时间
	ContractAddr string // 合约地址
	MethodSig    string // 方法签名
	IsEnable     string // 启用
	GasUsed      string // gas
	ChainId      string // gas
	Desc         string // 描述
}

// relayeradminSpecifiedGasColumns holds the columns for table specifiedGas.
var relayeradminSpecifiedGasColumns = RelayeradminSpecifiedGasColumns{
	Id:           "id",
	CreateTime:   "createTime",
	UpdateTime:   "updateTime",
	ContractAddr: "contractAddr",
	MethodSig:    "methodSig",
	IsEnable:     "isEnable",
	GasUsed:      "gasUsed",
	ChainId:      "chainId",
	Desc:         "desc",
}

// NewRelayeradminSpecifiedGasDao creates and returns a new DAO object for table data access.
func NewRelayeradminSpecifiedGasDao() *RelayeradminSpecifiedGasDao {
	return &RelayeradminSpecifiedGasDao{
		group:   "relayeradmin",
		table:   "specifiedGas",
		columns: relayeradminSpecifiedGasColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *RelayeradminSpecifiedGasDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *RelayeradminSpecifiedGasDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *RelayeradminSpecifiedGasDao) Columns() RelayeradminSpecifiedGasColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *RelayeradminSpecifiedGasDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *RelayeradminSpecifiedGasDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *RelayeradminSpecifiedGasDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
