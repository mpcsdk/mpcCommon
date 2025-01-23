// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// RelayeradminAssignFee is the golang structure of table assignFee for DAO operations like Where/Data.
type RelayeradminAssignFee struct {
	g.Meta     `orm:"table:assignFee, do:true"`
	Id         interface{} // ID
	CreateTime *gtime.Time // 创建时间
	UpdateTime *gtime.Time // 更新时间
	AppId      interface{} // 渠道号
	AppName    interface{} // 渠道名
	Fee        interface{} // fee
	IsEnable   interface{} // 启用
	ChainId    interface{} // gas
}
