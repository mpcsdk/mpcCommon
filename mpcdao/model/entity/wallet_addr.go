// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// WalletAddr is the golang structure for table wallet_addr.
type WalletAddr struct {
	UserId     string `json:"userId"     orm:"user_id"     ` //
	WalletAddr string `json:"walletAddr" orm:"wallet_addr" ` //
	ChainId    int64  `json:"chainId"    orm:"chain_id"    ` //
}
