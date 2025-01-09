// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// FcmOfflineMsg is the golang structure for table fcm_offline_msg.
type FcmOfflineMsg struct {
	FmcToken    string      `json:"fmcToken"    orm:"fmc_token"    ` //
	Title       string      `json:"title"       orm:"title"        ` //
	Body        string      `json:"body"        orm:"body"         ` //
	Data        string      `json:"data"        orm:"data"         ` //
	Address     string      `json:"address"     orm:"address"      ` //
	UserId      string      `json:"userId"      orm:"user_id"      ` //
	CreatedTime *gtime.Time `json:"createdTime" orm:"created_time" ` //
	Id          string      `json:"id"          orm:"id"           ` //
}
