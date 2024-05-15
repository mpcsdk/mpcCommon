// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PushErr is the golang structure for table push_err.
type PushErr struct {
	FmcToken    string      `json:"fmcToken"    ` //
	Title       string      `json:"title"       ` //
	Body        string      `json:"body"        ` //
	Data        string      `json:"data"        ` //
	Err         string      `json:"err"         ` //
	Address     string      `json:"address"     ` //
	UserId      string      `json:"userId"      ` //
	CreatedTime *gtime.Time `json:"createdTime" ` //
}
