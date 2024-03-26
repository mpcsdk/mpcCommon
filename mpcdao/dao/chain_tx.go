// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"github.com/mpcsdk/mpcCommon/mpcdao/dao/internal"
)

// internalChainTxDao is internal type for wrapping internal DAO implements.
type internalChainTxDao = *internal.ChainTxDao

// chainTxDao is the data access object for table chain_tx.
// You can define custom methods on it to extend its functionality as you wish.
type chainTxDao struct {
	internalChainTxDao
}

var (
	// ChainTx is globally public accessible object for table chain_tx operations.
	ChainTx = chainTxDao{
		internal.NewChainTxDao(),
	}
)

// Fill with you ideas below.