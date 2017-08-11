package randutil

import (
	"crypto/rand"
	"encoding/binary"
	"errors"
	"math"
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

func (i *cryptoIntner) Uint64() (uint64, error) {
	var u big.Int
	u.SetUint64(math.MaxUint64)
	v, err := rand.Int(rand.Reader, &u)
	if err != nil {
		return 0, err
	}
	return v.Uint64(), nil
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
