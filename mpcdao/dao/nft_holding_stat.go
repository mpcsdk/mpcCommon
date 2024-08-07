// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"github.com/mpcsdk/mpcCommon/mpcdao/dao/internal"
)

// internalNftHoldingStatDao is internal type for wrapping internal DAO implements.
type internalNftHoldingStatDao = *internal.NftHoldingStatDao

// nftHoldingStatDao is the data access object for table nft_holding_stat.
// You can define custom methods on it to extend its functionality as you wish.
type nftHoldingStatDao struct {
	internalNftHoldingStatDao
}

var (
	// NftHoldingStat is globally public accessible object for table nft_holding_stat operations.
	NftHoldingStat = nftHoldingStatDao{
		internal.NewNftHoldingStatDao(),
	}
)

// Fill with you ideas below.
