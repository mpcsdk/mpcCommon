package mpcdaoutil

import (
	"fmt"
)

func RiskadminContractabiKey(chainId int64, contractAddr string) string {
	return fmt.Sprintf("%d:%s", chainId, contractAddr)
}
func RelayerAdminAssignFeeKey(chainId int, appId string) string {
	return fmt.Sprintf("%d:%s", chainId, appId)
}
func RelayerAdminSpecifiedGas(chainId int, appId string) string {
	return fmt.Sprintf("%d:%s", chainId, appId)
}
