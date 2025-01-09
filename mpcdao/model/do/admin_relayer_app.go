// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminRelayerApp is the golang structure of table relayer_app for DAO operations like Where/Data.
type AdminRelayerApp struct {
	g.Meta     `orm:"table:relayer_app, do:true"`
	Id         interface{} // ID
	CreateTime *gtime.Time // 创建时间
	UpdateTime *gtime.Time // 更新时间
	AppId      interface{} // 渠道号
	AppName    interface{} // 渠道名
	IsEnable   interface{} // 启用
}
