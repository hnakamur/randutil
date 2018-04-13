package v2

import (
	"encoding/binary"
	"math"
)

const uint64ByteCount = 8

// Reader is a type for reading random bytes.
type Reader struct {
	intner Intner
	left   []byte
	buf    [uint64ByteCount]byte
}

// NewReader creates a new random reader.
func NewReader(i Intner) *Reader {
	return &Reader{intner: i}
}

// Read reads random bytes to the buffer.
// To make Reader usable as io.Reader, Read method has two return values,
// the number of bytes read and an error, but actually it never returns
// an error.
func (r *Reader) Read(p []byte) (int, error) {
	total := 0
	if len(r.left) > 0 {
		n := len(r.left)
		if len(p) < n {
			n = len(p)
		}
		copy(p, r.left[:n])
		r.left = r.left[n:]
		p = p[n:]
		total += n
	}
	for len(p) >= uint64ByteCount {
		v := r.intner.Int63n(math.MaxInt64)
		binary.BigEndian.PutUint64(p, uint64(v))
		p = p[uint64ByteCount:]
		total += uint64ByteCount
	}
	if len(p) > 0 {
		v := r.intner.Int63n(math.MaxInt64)
		r.left = r.buf[:]
		binary.BigEndian.PutUint64(r.left, uint64(v))
		copy(p, r.left[:len(p)])
		r.left = r.left[len(p):]
		total += len(p)
	}
	return total, nil
}
