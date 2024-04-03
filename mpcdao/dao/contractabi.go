// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"github.com/mpcsdk/mpcCommon/mpcdao/dao/internal"
)

// internalContractabiDao is internal type for wrapping internal DAO implements.
type internalContractabiDao = *internal.ContractabiDao

// contractabiDao is the data access object for table contractabi.
// You can define custom methods on it to extend its functionality as you wish.
type contractabiDao struct {
	internalContractabiDao
}

var (
	// Contractabi is globally public accessible object for table contractabi operations.
	Contractabi = contractabiDao{
		internal.NewContractabiDao(),
	}
)

// Fill with you ideas below.
