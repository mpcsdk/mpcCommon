// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// FcmToken is the golang structure for table fcm_token.
type FcmToken struct {
	UserId      string      `json:"userId"      ` //
	FcmToken    string      `json:"fcmToken"    ` //
	Token       string      `json:"token"       ` //
	Address     string      `json:"address"     ` //
	CreatedTime *gtime.Time `json:"createdTime" ` //
	UpdatedTime *gtime.Time `json:"updatedTime" ` //
}