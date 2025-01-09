// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"github.com/mpcsdk/mpcCommon/mpcdao/dao/internal"
)

// internalRelayerReceiptDao is internal type for wrapping internal DAO implements.
type internalRelayerReceiptDao = *internal.RelayerReceiptDao

// relayerReceiptDao is the data access object for table receipt.
// You can define custom methods on it to extend its functionality as you wish.
type relayerReceiptDao struct {
	internalRelayerReceiptDao
}

var (
	// RelayerReceipt is globally public accessible object for table receipt operations.
	RelayerReceipt = relayerReceiptDao{
		internal.NewRelayerReceiptDao(),
	}
)

// Fill with you ideas below.
