package randutil

import "math/rand"

type MathIntner struct {
	rand *rand.Rand
}

func NewMathIntner(seed int64) Intner {
	return &MathIntner{rand: rand.New(rand.NewSource(seed))}
}

func (i *MathIntner) Intn(n int) (int, error) {
	return i.rand.Intn(n), nil
}
