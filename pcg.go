package randutil

import (
	"math"

	"github.com/davidminor/gorand/pcg"
)

type pcg64Intner struct {
	rand pcg.Pcg64
}

type pcg32Intner struct {
	rand pcg.Pcg32
}

// NewPCG64Intner returns a new Intner using pcg.Pcg64.
func NewPCG64Intner(seedH, seedL uint64) Intner {
	return &pcg64Intner{rand: pcg.NewPcg64(seedH, seedL)}
}

func (i *pcg64Intner) Intn(n int) (int, error) {
	return int(i.rand.NextN(uint64(n))), nil
}

func (i *pcg64Intner) Int63n(n int64) (int64, error) {
	return int64(i.rand.NextN(uint64(n))), nil
}

// NewPCG32Intner returns a new Intner using pcg.Pcg32.
func NewPCG32Intner(seed uint64) Intner {
	return &pcg32Intner{rand: pcg.NewPcg32(seed)}
}

func (i *pcg32Intner) Intn(n int) (int, error) {
	if n <= math.MaxUint32 {
		return int(i.rand.NextN(uint32(n))), nil
	}
	h := int(i.rand.Next())
	l := int(i.rand.Next())
	return (h<<32 | l) % n, nil
}

func (i *pcg32Intner) Int63n(n int64) (int64, error) {
	if n <= math.MaxUint32 {
		return int64(i.rand.NextN(uint32(n))), nil
	}
	h := int64(i.rand.Next())
	l := int64(i.rand.Next())
	return (h<<32 | l) % n, nil
}
