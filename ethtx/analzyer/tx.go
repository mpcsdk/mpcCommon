package analzyer

import (
	"encoding/hex"
	"encoding/json"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/mpcsdk/mpcCommon/mpcmodel"
)

// /input
type SignTx struct {
	ChainId uint64         `json:"chainId,omitempty"`
	Address common.Address `json:"address,omitempty"`
	Number  uint64         `json:"number,omitempty"`
	Txs     []*SignTxData  `json:"txs,omitempty"`
	TxHash  string         `json:"txHash,omitempty"`
}
type SignTxData struct {
	Target common.Address `json:"target,omitempty"`
	Data   string         `json:"data,omitempty"`
}

// //output
type AnalzyedTx struct {
	Address string
	Txs     []*AnalzyedTxData
}
type AnalzyedTxData struct {
	Contract   string
	MethodId   string
	MethodName string
	MethodSig  string
	Data       string
	Args       map[string]interface{}
	////
	From  string
	To    string
	Value *big.Int
}

func (s *Analzyer) SignTx(signData string) (*SignTx, error) {
	signtx := &SignTx{}
	err := json.Unmarshal([]byte(signData), signtx)
	if err != nil {
		return nil, err
	}
	return signtx, nil
	///
	// atx := &AnalzyedTx{}
	// atx.Address = strings.ToLower(signtx.Address)
	// ///
	// for _, tx := range signtx.Txs {
	// 	adata, err := s.analzyTx(tx)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	if adata == nil {
	// 		continue
	// 	}
	// 	atx.Txs = append(atx.Txs, adata)
	// }
	// return atx, nil
}

func (s *Analzyer) AnalzyTxData(tx *SignTxData, contractRule *mpcmodel.ContractRule) (*AnalzyedTxData, error) {
	// tx.Target = strings.ToLower(tx.Target)
	if abistr, ok := s.abis[tx.Target.String()]; !ok {
		return nil, nil
		// return nil, errors.New("abi not found:" + tx.Target)
	} else {
		///
		contract, err := abi.JSON(strings.NewReader(abistr))
		if err != nil {
			return nil, err
		}
		//data
		dataByte, err := hex.DecodeString(strings.TrimPrefix(tx.Data, "0x"))
		if err != nil {
			return nil, err
		}
		////
		method, err := contract.MethodById(dataByte[:4])
		if err != nil {
			return nil, err
		}
		args := make(map[string]interface{})
		err = method.Inputs.UnpackIntoMap(args, dataByte[4:])
		if err != nil {
			return nil, err
		}
		from := ""
		to := ""
		val := big.NewInt(0)
		if v, ok := args[contractRule.MethodFromField]; ok {
			from = v.(common.Address).Hex()
		}
		if v, ok := args[contractRule.MethodToField]; ok {
			to = v.(common.Address).Hex()
		}
		if v, ok := args[contractRule.MethodValueField]; ok {
			val = v.(*big.Int)
		}
		atx := &AnalzyedTxData{
			Contract:   tx.Target.String(),
			MethodId:   hex.EncodeToString(method.ID),
			MethodName: method.RawName,
			MethodSig:  method.Sig,
			Data:       tx.Data,
			Args:       args,
			///
			From:  from,
			To:    to,
			Value: val,
		}
		return atx, nil
	}
}

// /
// func (s *Analzyer) AnalzyTxDataNFT(contract string, tx *SignTxData, nftrule *mpcmodel.ContractRule) (*AnalzyedTxData, error) {
// 	// tx.Target = strings.ToLower(tx.Target)
// 	if abistr, ok := s.abis[tx.Target.String()]; !ok {
// 		return nil, nil
// 	} else {
// 		///
// 		contract, err := abi.JSON(strings.NewReader(abistr))
// 		if err != nil {
// 			return nil, err
// 		}
// 		//data
// 		dataByte, err := hex.DecodeString(strings.TrimPrefix(tx.Data, "0x"))
// 		if err != nil {
// 			return nil, err
// 		}
// 		////
// 		method, err := contract.MethodById(dataByte[:4])
// 		if err != nil {
// 			return nil, err
// 		}
// 		args := make(map[string]interface{})
// 		err = method.Inputs.UnpackIntoMap(args, dataByte[4:])
// 		if err != nil {
// 			return nil, err
// 		}
// 		///
// 		from := ""
// 		to := ""
// 		val := big.NewInt(0)
// 		if v, ok := args[nftrule.MethodFromField]; ok {
// 			from = strings.ToLower(v.(common.Address).Hex())
// 		}
// 		if v, ok := args[nftrule.MethodToField]; ok {
// 			to = strings.ToLower(v.(common.Address).Hex())
// 		}
// 		if v, ok := args[nftrule.MethodValueField]; ok {
// 			val = v.(*big.Int)
// 		}
// 		atx := &AnalzyedTxData{
// 			Contract:   tx.Target.String(),
// 			MethodId:   hex.EncodeToString(method.ID),
// 			MethodName: method.RawName,
// 			MethodSig:  method.Sig,
// 			Data:       tx.Data,
// 			Args:       args,
// 			///
// 			From:  from,
// 			To:    to,
// 			Value: val,
// 		}
// 		return atx, nil
// 	}
// }

// func (s *Analzyer) AnalzyTxDataFT(contract string, tx *SignTxData, contractrule *mpcmodel.ContractRule) (*AnalzyedTxData, error) {
// 	// tx.Target = strings.ToLower(tx.Target)
// 	if abistr, ok := s.abis[tx.Target.String()]; !ok {
// 		return nil, nil
// 	} else {
// 		///
// 		contract, err := abi.JSON(strings.NewReader(abistr))
// 		if err != nil {
// 			return nil, err
// 		}
// 		//data
// 		dataByte, err := hex.DecodeString(strings.TrimPrefix(tx.Data, "0x"))
// 		if err != nil {
// 			return nil, err
// 		}
// 		////
// 		method, err := contract.MethodById(dataByte[:4])
// 		if err != nil {
// 			return nil, err
// 		}
// 		args := make(map[string]interface{})
// 		err = method.Inputs.UnpackIntoMap(args, dataByte[4:])
// 		if err != nil {
// 			return nil, err
// 		}
// 		///
// 		from := ""
// 		to := ""
// 		val := big.NewInt(0)
// 		if v, ok := args[ftrule.MethodFromField]; ok {
// 			from = strings.ToLower(v.(common.Address).Hex())
// 		}
// 		if v, ok := args[ftrule.MethodToField]; ok {
// 			to = strings.ToLower(v.(common.Address).Hex())
// 		}
// 		if v, ok := args[ftrule.MethodValueField]; ok {
// 			val = v.(*big.Int)
// 		}
// 		atx := &AnalzyedTxData{
// 			Contract:   tx.Target.String(),
// 			MethodId:   hex.EncodeToString(method.ID),
// 			MethodName: method.RawName,
// 			MethodSig:  method.Sig,
// 			Data:       tx.Data,
// 			Args:       args,
// 			///
// 			From:  from,
// 			To:    to,
// 			Value: val,
// 		}
// 		return atx, nil
// 	}
// }
