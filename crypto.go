package randutil

import (
	"crypto/rand"
	"math/big"
)

type CryptoIntner struct{}

func NewCryptoIntner() Intner {
	return new(CryptoIntner)
}

func (i *CryptoIntner) Intn(n int) (int, error) {
	v, err := rand.Int(rand.Reader, big.NewInt(int64(n)))
	if err != nil {
		return 0, err
	}
	return int(v.Int64()), nil
}
