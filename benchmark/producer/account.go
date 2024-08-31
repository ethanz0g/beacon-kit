package producer

import "github.com/ethereum/go-ethereum/crypto"

type Account struct {
	Index      uint32
	Nonce      uint32
	Address    string
	Checksum   string
	PrivateKey []byte
	IsFaucet   bool
}

func NewAccount(index uint32) Account {
	pk := generateEthSecp256k1PrivateKeyByUint32(index)
	addr := crypto.PubkeyToAddress(pk.PublicKey).Hex()
	return Account{
		Index:      index,
		Nonce:      0,
		Address:    addr,
		Checksum:   toChecksumAddress(addr),
		PrivateKey: crypto.FromECDSA(pk),
	}
}

func CreateFaucetAccount(privateKey string) Account {
	// get private key from genesis file
	pk := loadPrivateKey(privateKey)
	addr := crypto.PubkeyToAddress(pk.PublicKey).Hex()
	return Account{
		IsFaucet:   true,
		Index:      0,
		Nonce:      0,
		Address:    addr,
		Checksum:   toChecksumAddress(addr),
		PrivateKey: crypto.FromECDSA(pk),
	}
}

func (a *Account) GetAndIncrementNonce() uint32 {
	now := a.Nonce
	a.Nonce += 1
	return now
}
