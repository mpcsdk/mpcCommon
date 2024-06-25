// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PushErr is the golang structure for table push_err.
type PushErr struct {
	FmcToken    string      `json:"fmcToken"    orm:"fmc_token"    ` //
	Title       string      `json:"title"       orm:"title"        ` //
	Body        string      `json:"body"        orm:"body"         ` //
	Data        string      `json:"data"        orm:"data"         ` //
	Err         string      `json:"err"         orm:"err"          ` //
	Address     string      `json:"address"     orm:"address"      ` //
	UserId      string      `json:"userId"      orm:"user_id"      ` //
	CreatedTime *gtime.Time `json:"createdTime" orm:"created_time" ` //
}
