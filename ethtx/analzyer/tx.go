package analzyer

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

// /input
type SignData struct {
	ChainId uint64         `json:"chainId,omitempty"`
	Address common.Address `json:"address,omitempty"`
	Number  uint64         `json:"number,omitempty"`
	Txs     []*SignTx      `json:"txs,omitempty"`
	TxHash  string         `json:"txHash,omitempty"`
}
type BigNumberish struct {
	Type   string `json:"type"`
	Hex    string `json:"hex"`
	bigint *big.Int
}

func (s *BigNumberish) BigInt() *big.Int {
	if s.bigint == nil {
		if s.Type == "BigNumber" {
			b := big.NewInt(0)
			b.UnmarshalText([]byte(s.Hex))
			s.bigint = b
		}
	}
	return s.bigint
}

type SignTx struct {
	Target common.Address `json:"target,omitempty"`
	Value  *BigNumberish  `json:"value,omitempty"`
	Data   string         `json:"data,omitempty"`
}

// //output
type AnalzyedSignData struct {
	Target string
	Txs    []*AnalzyedSignTx
}

type AnalzyedSignTx struct {
	Target     string
	MethodId   string
	MethodName string
	MethodSig  string
	Data       string
	Args       map[string]interface{}
	////method
	From     string
	To       string
	Value    *BigInt
	TokenId  *BigInt
	Values   []*big.Int
	TokenIds []*big.Int
}

func DeSignData(signDataStr string) (*SignData, error) {
	signdata := &SignData{}
	err := json.Unmarshal([]byte(signDataStr), signdata)
	if err != nil {
		return nil, err
	}
	return signdata, nil
	///
}

//	func (s *Analzyer) AnalzySignData(signData *SignData) (*AnalzyedSignData, error) {
//		chainAbi := s.chainAbi[signData.ChainId]
//		if chainAbi == nil {
//			return nil, mpccode.CodeParamInvalid()
//		}
//		/////
//		asdata := &AnalzyedSignData{
//			Address: signData.Address.Hex(),
//			Txs:     []*AnalzyedSignTx{},
//		}
//		for _, tx := range signData.Txs {
//			atx, err := s.AnalzySignTx(signData.ChainId, tx, crule)
//			if err != nil {
//				return nil, mpccode.CodeParamInvalid()
//			}
//			asdata.Txs = append(asdata.Txs, atx)
//		}
//		return asdata, nil
//	}

func (s *Analzyer) AnalzySignTx(signTx *SignTx) (*AnalzyedSignTx, error) {
	///////
	contractAbi := s.abis[signTx.Target.Hex()]
	if contractAbi == nil {
		return nil, errors.New("unsupport contract:" + signTx.Target.String())
	}
	//data
	dataByte, err := hex.DecodeString(strings.TrimPrefix(signTx.Data, "0x"))
	if err != nil {
		return nil, err
	}
	////
	method, err := contractAbi.abi.MethodById(dataByte[:4])
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
	tokenId := big.NewInt(0)
	ids := []*big.Int{}
	vals := []*big.Int{}
	//todo:
	if contractAbi.kind == "erc20" {
		if method.Name != "transfer" {
			return nil, errors.New("unsupport contract method:" + method.Name)
		}

		if v, ok := args["recipient"]; ok {
			to = v.(common.Address).Hex()
		}
		if v, ok := args["amount"]; ok {
			val = v.(*big.Int)
		}
	} else if contractAbi.kind == "erc721" {
		if method.Name != "transferFrom" {
			return nil, errors.New("unsupport contract method:" + method.Name)
		}
		if v, ok := args["from"]; ok {
			from = v.(common.Address).Hex()
		}
		if v, ok := args["to"]; ok {
			to = v.(common.Address).Hex()
		}
		if v, ok := args["tokenId"]; ok {
			tokenId = v.(*big.Int)
		}
	} else if contractAbi.kind == "erc1155" {
		if method.Name != "safeTransferFrom" {
			return nil, errors.New("unsupport contract method:" + method.Name)
		}
		if v, ok := args["from"]; ok {
			from = v.(common.Address).Hex()
		}
		if v, ok := args["to"]; ok {
			to = v.(common.Address).Hex()
		}

		if v, ok := args["id"]; ok {
			tokenId = v.(*big.Int)
		}
		if v, ok := args["value"]; ok {
			val = v.(*big.Int)
		}

	}

	atx := &AnalzyedSignTx{
		Target:     signTx.Target.String(),
		MethodId:   hex.EncodeToString(method.ID),
		MethodName: method.RawName,
		MethodSig:  method.Sig,
		Data:       signTx.Data,
		Args:       args,
		///
		From:     from,
		To:       to,
		Value:    (*BigInt)(val),
		TokenId:  (*BigInt)(tokenId),
		TokenIds: ids,
		Values:   vals,
	}
	return atx, nil
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
