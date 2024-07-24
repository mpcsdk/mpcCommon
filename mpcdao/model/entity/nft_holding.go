// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// NftHolding is the golang structure for table nft_holding.
type NftHolding struct {
	ChainId     int64       `json:"chainId"     orm:"chain_id"     ` //
	Address     string      `json:"address"     orm:"address"      ` //
	Contract    string      `json:"contract"    orm:"contract"     ` //
	TokenId     string      `json:"tokenId"     orm:"token_id"     ` //
	Value       int64       `json:"value"       orm:"value"        ` //
	BlockNumber int64       `json:"blockNumber" orm:"block_number" ` //
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"   ` //
	Kind        string      `json:"kind"        orm:"kind"         ` //
}
