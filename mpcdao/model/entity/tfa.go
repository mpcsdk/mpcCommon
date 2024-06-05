// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Tfa is the golang structure for table tfa.
type Tfa struct {
	UserId         string      `json:"userId"         orm:"user_id"          ` //
	CreatedAt      *gtime.Time `json:"createdAt"      orm:"created_at"       ` //
	DeletedAt      *gtime.Time `json:"deletedAt"      orm:"deleted_at"       ` //
	Phone          string      `json:"phone"          orm:"phone"            ` //
	Mail           string      `json:"mail"           orm:"mail"             ` //
	PhoneUpdatedAt *gtime.Time `json:"phoneUpdatedAt" orm:"phone_updated_at" ` //
	MailUpdatedAt  *gtime.Time `json:"mailUpdatedAt"  orm:"mail_updated_at"  ` //
	TokenData      string      `json:"tokenData"      orm:"token_data"       ` //
	TxNeedVerify   bool        `json:"txNeedVerify"   orm:"tx_need_verify"   ` //
}
