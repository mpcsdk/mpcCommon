// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// MpcContext is the golang structure for table mpc_context.
type MpcContext struct {
	UserId    string      `json:"userId"    orm:"user_id"    ` //
	Context   string      `json:"context"   orm:"context"    ` //
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" ` //
	Request   string      `json:"request"   orm:"request"    ` //
	Token     string      `json:"token"     orm:"token"      ` //
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" ` //
	DeletedAt *gtime.Time `json:"deletedAt" orm:"deleted_at" ` //
	PubKey    string      `json:"pubKey"    orm:"pub_key"    ` //
	TokenData string      `json:"tokenData" orm:"token_data" ` //
}
