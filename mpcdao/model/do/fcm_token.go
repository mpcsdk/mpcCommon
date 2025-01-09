// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// FcmToken is the golang structure of table fcm_token for DAO operations like Where/Data.
type FcmToken struct {
	g.Meta      `orm:"table:fcm_token, do:true"`
	UserId      interface{} // 用户id
	FcmToken    interface{} //
	Token       interface{} //
	Address     interface{} //
	CreatedTime *gtime.Time //
	UpdatedTime *gtime.Time //
}
