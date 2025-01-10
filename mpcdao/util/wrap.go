package mpcdaoutil

import (
	"fmt"
)

func RiskadminContractabiKey(chainId int64, contractAddr string) string {
	return fmt.Sprintf("%d:%s", chainId, contractAddr)
}
