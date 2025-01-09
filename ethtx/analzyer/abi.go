package analzyer

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

type abiStruct struct {
	abistr  string
	abi     *abi.ABI
	kind    string
	decimal int
	// abiStructs map[string]*abi.ABI
}

type Analzyer struct {
	abis map[string]*abiStruct
}

func NewAnalzer() *Analzyer {
	return &Analzyer{
		abis: map[string]*abiStruct{},
		// abis:       map[string]string{},
		// abiStructs: map[string]*abi.ABI{},
	}
}

func (s *Analzyer) AddAbi(contract string, abistr string, kind string, decimal int) error {
	abiabi, err := abi.JSON(strings.NewReader(abistr))
	if err != nil {
		return err
	}

	s.abis[contract] = &abiStruct{
		abistr:  abistr,
		abi:     &abiabi,
		kind:    kind,
		decimal: decimal,
	}

	return nil
}

// //
