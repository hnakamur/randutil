package randutil

import (
	"errors"
	"math"
	"unsafe"
)

// Encoder generates random values and encodes them in the specified digit characters.
type Encoder struct {
	intner Intner
	digits []byte
}

// NewEncoder creates an Encoder.
// The length of digits must be between 1 and 256.
func NewEncoder(intner Intner, digits []byte) (*Encoder, error) {
	l := len(digits)
	if l == 0 {
		return nil, errors.New("empty digits")
	}
	if l > 256 {
		return nil, errors.New("too many digits")
	}
	return &Encoder{
		intner: intner,
		digits: digits,
	}, nil
}

// Bytes writes random encoded bytes to the passed buffer.
func (e *Encoder) RandomBytes(buf []byte) error {
	// This implementation is inspired by math/rand.rand()
	// https://github.com/golang/go/blob/go1.9.3/src/math/rand/rand.go#L215-L230
	pos := 0
	var val int64
	for i := 0; i < len(buf); i++ {
		if pos == 0 {
			var err error
			val, err = e.intner.Int63n(math.MaxInt64)
			if err != nil {
				return err
			}
			pos = 7
		}
		buf[i] = e.digits[val%int64(len(e.digits))]
		val >>= 8
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
	// Convert byte buffer to string without an extra allocation, in the same way as
	// https://github.com/golang/go/blob/go1.10beta1/src/strings/builder.go#L22
	return *(*string)(unsafe.Pointer(&buf)), nil
}
