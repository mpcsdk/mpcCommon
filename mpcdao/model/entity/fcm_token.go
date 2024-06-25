// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// FcmToken is the golang structure for table fcm_token.
type FcmToken struct {
	UserId      string      `json:"userId"      orm:"user_id"      ` //
	FcmToken    string      `json:"fcmToken"    orm:"fcm_token"    ` //
	Token       string      `json:"token"       orm:"token"        ` //
	Address     string      `json:"address"     orm:"address"      ` //
	CreatedTime *gtime.Time `json:"createdTime" orm:"created_time" ` //
	UpdatedTime *gtime.Time `json:"updatedTime" orm:"updated_time" ` //
}
