// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// FcmOfflineMsg is the golang structure for table fcm_offline_msg.
type FcmOfflineMsg struct {
	FmcToken    string      `json:"fmcToken"    ` //
	Title       string      `json:"title"       ` //
	Body        string      `json:"body"        ` //
	Data        string      `json:"data"        ` //
	Address     string      `json:"address"     ` //
	UserId      string      `json:"userId"      ` //
	CreatedTime *gtime.Time `json:"createdTime" ` //
	Id          string      `json:"id"          ` //
}
