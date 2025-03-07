// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"github.com/mpcsdk/mpcCommon/mpcdao/dao/internal"
)

// internalRelayeradminSpecifiedGasDao is internal type for wrapping internal DAO implements.
type internalRelayeradminSpecifiedGasDao = *internal.RelayeradminSpecifiedGasDao

// relayeradminSpecifiedGasDao is the data access object for table specifiedGas.
// You can define custom methods on it to extend its functionality as you wish.
type relayeradminSpecifiedGasDao struct {
	internalRelayeradminSpecifiedGasDao
}

var (
	// RelayeradminSpecifiedGas is globally public accessible object for table specifiedGas operations.
	RelayeradminSpecifiedGas = relayeradminSpecifiedGasDao{
		internal.NewRelayeradminSpecifiedGasDao(),
	}
)

// Fill with you ideas below.
