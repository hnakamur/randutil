package randutil

import (
	"golang.org/x/exp/rand"
)

// Uniform is a uniform random chooser.
type Uniform struct {
	rnd *rand.Rand
}

// NewUniform creates a Uniform instance.
func NewUniform(src rand.Source) *Uniform {
	if src == nil {
		panic("src must not be nil")
	}
	return &Uniform{rnd: rand.New(src)}
}

// Choose chooses one value from values randomly.
func (u *Uniform) Choose(values []string) string {
	i := u.rnd.Intn(len(values))
	return values[i]
}
