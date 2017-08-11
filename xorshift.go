package randutil

import "github.com/lazybeaver/xorshift"

type xorShiftIntner struct {
	rand xorshift.XorShift
}

// NewXorShift64StarIntner returns a new Intner using xorshift.XorShift64Star.
func NewXorShift64StarIntner(seed uint64) Intner {
	return &xorShiftIntner{rand: xorshift.NewXorShift64Star(seed)}
}

// NewXorShift128PlusIntner returns a new Intner using xorshift.XorShift128Plus.
func NewXorShift128PlusIntner(seed uint64) Intner {
	return &xorShiftIntner{rand: xorshift.NewXorShift128Plus(seed)}
}

// NewXorShift1024StarIntner returns a new Intner using xorshift.XorShift1024Star.
func NewXorShift1024StarIntner(seed uint64) Intner {
	return &xorShiftIntner{rand: xorshift.NewXorShift1024Star(seed)}
}

func (i *xorShiftIntner) Intn(n int) (int, error) {
	return int(i.rand.Next() % uint64(n)), nil
}

func (i *xorShiftIntner) Int63n(n int64) (int64, error) {
	return int64(i.rand.Next() % uint64(n)), nil
}
