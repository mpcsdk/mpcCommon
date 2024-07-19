// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// NftHoldingStat is the golang structure for table nft_holding_stat.
type NftHoldingStat struct {
	ChainId     int64       `json:"chainId"     orm:"chain_id"     ` //
	BlockNumber int64       `json:"blockNumber" orm:"block_number" ` //
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"   ` //
}
