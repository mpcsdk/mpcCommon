// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"github.com/mpcsdk/mpcCommon/mpcdao/dao/internal"
)

// internalRelayeradminAppCfgDao is internal type for wrapping internal DAO implements.
type internalRelayeradminAppCfgDao = *internal.RelayeradminAppCfgDao

// relayeradminAppCfgDao is the data access object for table appCfg.
// You can define custom methods on it to extend its functionality as you wish.
type relayeradminAppCfgDao struct {
	internalRelayeradminAppCfgDao
}

var (
	// RelayeradminAppCfg is globally public accessible object for table appCfg operations.
	RelayeradminAppCfg = relayeradminAppCfgDao{
		internal.NewRelayeradminAppCfgDao(),
	}
)

// Fill with you ideas below.
