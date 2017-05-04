package randutil

import "math/rand"

type mathIntner struct {
	rand *rand.Rand
}

// NewMathIntner returns a new Intner using math/rand.
// This Intner is not safe for concurrent use by multiple goroutines.
func NewMathIntner(seed int64) Intner {
	return &mathIntner{rand: rand.New(rand.NewSource(seed))}
}

func (i *mathIntner) Intn(n int) (int, error) {
	return i.rand.Intn(n), nil
}
