package analzyer

import (
	"math/big"

	"github.com/shopspring/decimal"
)

type BigInt big.Int

func (s *BigInt) Int() *big.Int {
	return (*big.Int)(s)
}
func (s *BigInt) NewDecimal(value float64, exp int32) *BigInt {
	dec := decimal.NewFromFloat(value)
	dec = dec.Shift(exp)
	data := dec.BigInt()

	return (*BigInt)(data)
}
func (s *BigInt) NewInt(a int64) *BigInt {
	return (*BigInt)(big.NewInt(a))
}
func (s *BigInt) CmpInt(a int64) int {
	return s.Int().Cmp(big.NewInt(a))
}
func (s *BigInt) Cmp(a *BigInt) int {
	return s.Int().Cmp(a.Int())
}
func (s *BigInt) Add(a *BigInt) *BigInt {
	s.Int().Add(s.Int(), a.Int())
	return s
}
