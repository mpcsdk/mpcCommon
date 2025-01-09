package ethutil

import (
	"math"
	"math/big"
)

func BigDecimal2Float64(i *big.Int, decimal int) float64 {
	fbalance := big.NewFloat(0)
	fbalance.SetInt(i)
	fval := fbalance.Quo(fbalance, big.NewFloat(math.Pow10(decimal)))
	f, _ := fval.Float64()
	return f
}
