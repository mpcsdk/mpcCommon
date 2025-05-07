package keyutile

import (
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
)

type KeyStore struct {
	key *keystore.Key
}

func NewKeyStore() *KeyStore {
	return &KeyStore{}
}

func (s *KeyStore) Open(path string, passwd string) error {
	keyJSON, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	//
	key, err := keystore.DecryptKey(keyJSON, passwd)
	if err != nil {
		return err
	}
	s.key = key

	return nil
}

func (s *KeyStore) Key() *keystore.Key {
	return s.key
}

// func (s *KeyStore) Seed() ([]byte, error) {
// 	seed, err := generateSeedFromPrivateKey(s.key.PrivateKey)
// 	if err != nil {
// 		fmt.Println("err:", err)
// 		return nil, err
// 	}

// 	return seed, nil
// }
