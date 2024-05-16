// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// FcmOfflineMsg is the golang structure of table fcm_offline_msg for DAO operations like Where/Data.
type FcmOfflineMsg struct {
	g.Meta      `orm:"table:fcm_offline_msg, do:true"`
	FmcToken    interface{} //
	Title       interface{} //
	Body        interface{} //
	Data        interface{} //
	Err         interface{} //
	Address     interface{} //
	UserId      interface{} //
	CreatedTime *gtime.Time //
	Id          interface{} //
}
