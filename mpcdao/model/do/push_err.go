// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PushErr is the golang structure of table push_err for DAO operations like Where/Data.
type PushErr struct {
	g.Meta      `orm:"table:push_err, do:true"`
	FmcToken    interface{} //
	Title       interface{} //
	Body        interface{} //
	Data        interface{} //
	Err         interface{} //
	Address     interface{} //
	UserId      interface{} //
	CreatedTime *gtime.Time //
}
