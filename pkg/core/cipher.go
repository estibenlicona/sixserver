package core

import (
	"encoding/hex"
	"golang.org/x/crypto/blowfish"
	"log"
	"sixserver/pkg/types"
)

func NewCipher(hexKey string) *types.Cipher {

	key, err := hex.DecodeString(hexKey)
	if err != nil {
		log.Printf("Error decoding cipher key: %v\n", err)
	}

	block, err := blowfish.NewCipher(key)
	if err != nil {
		log.Printf("Error creating cipher block: %v\n", err)
	}

	return &types.Cipher{Block: block}
}
