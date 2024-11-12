// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// RelayerChannelFee is the golang structure for table relayer_channel_fee.
type RelayerChannelFee struct {
	AppId     string      `json:"appId"     orm:"app_id"    ` //
	Fee       float64     `json:"fee"       orm:"fee"       ` //
	Updatedat *gtime.Time `json:"updatedat" orm:"updatedat" ` //
}
