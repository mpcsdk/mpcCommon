// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SyncchainState is the golang structure for table syncchain_state.
type SyncchainState struct {
	ChainId      int64       `json:"chainId"      orm:"chain_id"      ` //
	CurrentBlock int64       `json:"currentBlock" orm:"current_block" ` //
	Createdat    *gtime.Time `json:"createdat"    orm:"createdat"     ` //
	Updatedat    *gtime.Time `json:"updatedat"    orm:"updatedat"     ` //
}
