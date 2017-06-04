package randutil

import (
	"crypto/rand"
	"fmt"
	"io"
	"math/big"
)

const base = 36
const randomIDEncodedByteLen = 25

// AppendRandomID appends an ID string encoded in base36
// (numerical digit and lower ASCII characters)
// of a 128-bit random unsigned bytes
// to a byte buffer.
// The ID is left zero padded to have the fixed length of 25.
//
// If randomReader is nil, cyrpto/rand.Reader is used.
func AppendRandomID(buf []byte, randomReader io.Reader) ([]byte, error) {
	var rbuf [randomIDEncodedByteLen]byte
	if randomReader == nil {
		randomReader = rand.Reader
	}
	_, err := randomReader.Read(rbuf[:16])
	if err != nil {
		return nil, fmt.Errorf("failed to generate random 128-bit unsigned bytes, err=%v", err)
	}

	p := new(big.Int).SetBytes(rbuf[:16]).Append(rbuf[:0], base)
	for i := randomIDEncodedByteLen - len(p) - 1; i >= 0; i-- {
		buf = append(buf, '0')
	}
	return append(buf, p...), nil
}

// RandomID return an ID string encoded in base36
// (numerical digit and lower ASCII characters)
// of a 128-bit random unsigned bytes.
// The ID is left zero padded to have the fixed length of 25.
func RandomID(randomReader io.Reader) (string, error) {
	var buf [randomIDEncodedByteLen]byte
	p, err := AppendRandomID(buf[:0], randomReader)
	if err != nil {
		return "", err
	}
	return string(p), nil
}
