// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// RelayerAppFee is the golang structure of table app_fee for DAO operations like Where/Data.
type RelayerAppFee struct {
	g.Meta    `orm:"table:app_fee, do:true"`
	AppId     interface{} //
	Fee       interface{} //
	Updatedat *gtime.Time //
}
