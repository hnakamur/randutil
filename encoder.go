package randutil

import (
	"errors"
	"math"
)

// Encoder generates random values and encodes them in the specified digit characters.
type Encoder struct {
	intner    Intner
	digits    []byte
	n         int64
	batchSize int
}

// NewEncoder creates an Encoder.
// The length of digits must be between 1 and 256.
func NewEncoder(intner Intner, digits []byte) (*Encoder, error) {
	l := int64(len(digits))
	if l == 0 {
		return nil, errors.New("empty digits")
	}
	if l > 256 {
		return nil, errors.New("too many digits")
	}

	batchSize := int(math.Log2(math.MaxUint64) / math.Log2(float64(l)))
	n := int64(1)
	for j := 0; j < batchSize; j++ {
		n *= l
	}
	return &Encoder{
		intner:    intner,
		digits:    digits,
		n:         n,
		batchSize: batchSize,
	}, nil
}

// Bytes writes random encoded bytes to the passed buffer.
func (e *Encoder) RandomBytes(buf []byte) error {
	// This implementation is inspired by math/rand.rand()
	// https://github.com/golang/go/blob/go1.9.3/src/math/rand/rand.go#L215-L230
	pos := 0
	var val int64
	l := int64(len(e.digits))
	n := e.n
	for i := 0; i < len(buf); i++ {
		if pos == 0 {
			var err error
			val, err = e.intner.Int63n(n)
			if err != nil {
				return err
			}
			pos = batchSize - 1
		}
		buf[i] = e.digits[val%l]
		val /= l
		pos--
	}
	return nil
}

// String returns a random encoded string of the specified length.
func (e *Encoder) RandomString(length int) (string, error) {
	buf := make([]byte, length)
	err := e.RandomBytes(buf)
	if err != nil {
		return "", err
	}
	return string(buf), nil
}
