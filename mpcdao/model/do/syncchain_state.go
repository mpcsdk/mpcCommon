// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SyncchainState is the golang structure of table state for DAO operations like Where/Data.
type SyncchainState struct {
	g.Meta       `orm:"table:state, do:true"`
	ChainId      interface{} //
	CurrentBlock interface{} //
	Createdat    *gtime.Time //
	Updatedat    *gtime.Time //
}
