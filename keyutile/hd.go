package keyutile

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
)

type HDWallet struct {
	keyList []uint64
	keyMap  map[uint64]*ecdsa.PrivateKey
}

func generateSeedFromPrivateKey(privateKey *ecdsa.PrivateKey) ([]byte, error) {
	entropy := crypto.FromECDSA(privateKey)
	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		return nil, err
	}
	seed := bip39.NewSeed(mnemonic, "")
	return seed, nil
}
func derivePrivateKeyByInt64Index(masterKey *bip32.Key, index uint64) (*ecdsa.PrivateKey, error) {
	high32 := uint32(index >> 32)
	low32 := uint32(index)

	intermediateKey, err := masterKey.NewChildKey(high32)
	if err != nil {
		return nil, err
	}

	childKey, err := intermediateKey.NewChildKey(low32)
	if err != nil {
		return nil, err
	}

	privateKey, err := crypto.ToECDSA(childKey.Key)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}
func NewHDWallet(privateKey *ecdsa.PrivateKey, keyList []uint64) (*HDWallet, error) {
	// seed, err := generateSeedFromPrivateKey(privateKey)
	// if err != nil {
	// 	return nil, err
	// }
	// fmt.Println("seed:", common.Bytes2Hex(seed))
	seed := privateKey.D.Bytes()
	// fmt.Println("seed:", common.Bytes2Hex(seed))

	masterKey, err := bip32.NewMasterKey(seed)
	if err != nil {
		return nil, err
	}
	///
	wallet := &HDWallet{
		keyList: keyList,
		keyMap:  make(map[uint64]*ecdsa.PrivateKey),
	}
	////
	for _, i := range keyList {
		childKey, err := derivePrivateKeyByInt64Index(masterKey, i)
		if err != nil {
			return nil, err
		}
		wallet.keyMap[i] = childKey
	}
	return wallet, nil
}

func (s *HDWallet) Show() {
	for _, i := range s.keyList {
		fmt.Println("index:", i)
		fmt.Println("privateKey:", common.Bytes2Hex(crypto.FromECDSA(s.keyMap[i])))
		fmt.Println("publicKey:", common.Bytes2Hex(crypto.FromECDSAPub(&s.keyMap[i].PublicKey)))
		fmt.Println("address:", crypto.PubkeyToAddress(s.keyMap[i].PublicKey).Hex())
	}
}
func (s *HDWallet) KeyIdx(idx uint64) *ecdsa.PrivateKey {
	return s.keyMap[idx]
}
func (s *HDWallet) Keys() map[uint64]*ecdsa.PrivateKey {
	return s.keyMap
}
