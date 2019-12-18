package randutil

import (
	crand "crypto/rand"
	"encoding/binary"
	"time"
)

// NewSeed creats a new random seed.
func NewSeed() uint64 {
	var b [8]byte
	if _, err := crand.Read(b[:]); err != nil {
		return uint64(time.Now().UnixNano())
	}
	return binary.BigEndian.Uint64(b[:])
}
