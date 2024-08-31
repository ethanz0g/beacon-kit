package producer

import (
	"bytes"
	"encoding/hex"
	"strings"
	"time"

	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
	"golang.org/x/exp/rand"
)

func generateEthSecp256k1PrivateKeyByUint32(n uint32) *ecdsa.PrivateKey {
	privKeyInt := new(big.Int).SetUint64(uint64(n))
	curve := crypto.S256()
	privateKey, err := ecdsa.GenerateKey(curve, bytes.NewReader(privKeyInt.Bytes()))
	if err != nil {
		panic(err)
	}

	return privateKey
}

func toChecksumAddress(address string) string {
	address = strings.ToLower(strings.TrimPrefix(address, "0x"))

	hasher := sha3.NewLegacyKeccak256()
	hasher.Write([]byte(address))
	hash := hasher.Sum(nil)
	hashHex := hex.EncodeToString(hash)

	checksummedAddress := "0x"
	for i, c := range address {
		if c >= '0' && c <= '9' {
			checksummedAddress += string(c)
		} else {
			if hashHex[i] >= '8' {
				checksummedAddress += strings.ToUpper(string(c))
			} else {
				checksummedAddress += string(c)
			}
		}
	}

	return checksummedAddress
}

func loadPrivateKey(key string) *ecdsa.PrivateKey {
	privateKeyBytes, err := hex.DecodeString(key)
	if err != nil {
		panic(err)
	}

	privateKey, err := crypto.ToECDSA(privateKeyBytes)
	if err != nil {
		panic(err)
	}
	return privateKey
}

func shuffle(slice []uint32) []uint32 {
	rand.Seed(uint64(time.Now().UnixNano()))
	n := len(slice)
	for i := n - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}
