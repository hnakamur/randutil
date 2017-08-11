package randutil

import (
	"crypto/rand"
	"encoding/binary"
	"errors"
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

func (i *cryptoIntner) Int63n(n int64) (int64, error) {
	v, err := rand.Int(rand.Reader, big.NewInt(n))
	if err != nil {
		return 0, err
	}
	return v.Int64(), nil
}

// RandomSeed returns a random seed using the cryptographically
// secure pseudo random number generator.
func RandomSeed() (int64, error) {
	var buf [8]byte
	n, err := rand.Read(buf[:])
	if err != nil {
		return 0, err
	}
	if n != 8 {
		return 0, errors.New("got unexpected length of random bytes")
	}
	return int64(binary.LittleEndian.Uint64(buf[:])), nil
}
