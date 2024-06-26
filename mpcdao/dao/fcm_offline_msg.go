// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"github.com/mpcsdk/mpcCommon/mpcdao/dao/internal"
)

// internalFcmOfflineMsgDao is internal type for wrapping internal DAO implements.
type internalFcmOfflineMsgDao = *internal.FcmOfflineMsgDao

// fcmOfflineMsgDao is the data access object for table fcm_offline_msg.
// You can define custom methods on it to extend its functionality as you wish.
type fcmOfflineMsgDao struct {
	internalFcmOfflineMsgDao
}

var (
	// FcmOfflineMsg is globally public accessible object for table fcm_offline_msg operations.
	FcmOfflineMsg = fcmOfflineMsgDao{
		internal.NewFcmOfflineMsgDao(),
	}
)

// Fill with you ideas below.
