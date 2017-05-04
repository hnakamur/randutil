package randutil

import (
	"crypto/rand"
	"math/big"
)

type cryptoIntner struct{}

// NewCryptoIntner returns a new Intner using crypto/rand.
func NewCryptoIntner() Intner {
	return new(cryptoIntner)
}

func (i *cryptoIntner) Intn(n int) (int, error) {
	v, err := rand.Int(rand.Reader, big.NewInt(int64(n)))
	if err != nil {
		return 0, err
	}
	return int(v.Int64()), nil
}
