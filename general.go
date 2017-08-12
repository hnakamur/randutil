package randutil

import (
	"errors"
	"fmt"
)

// General is a general discrete distributed random value generator.
type General struct {
	intner Intner
	g      []float64
}

// NewGeneral creates a general discrete distributed random value generator.
//
// The rules for g are:
// * The lenght of g must be greater than or equal to 2.
// * The first element of g must be 0.
// * The last element of g must be 1.
// * The g must be monotonously increasing, that is, g[i] >= g[i-1] for any i.
func NewGeneral(intner Intner, g []float64) (*General, error) {
	if len(g) < 2 {
		return nil, errors.New("length of g must be greater than or equal to 2")
	}
	if g[0] != 0 {
		return nil, errors.New("the first element of g must be 0")
	}
	if g[len(g)-1] != 1 {
		return nil, errors.New("the last element of g must be 1")
	}
	for i := 1; i < len(g)-1; i++ {
		if g[i] < g[i-1] {
			return nil, fmt.Errorf("g must be monotonously increasing, g[%d]=%f < g[%d]=%f", i, g[i], i-1, g[i-1])
		}
	}
	return &General{intner: intner, g: g}, nil
}

// Next returns a general discrete distributed random value.
func (g *General) Next() (float64, error) {
	// http://www.ishikawa-lab.com/montecarlo/4shou.html
	gs := g.g
	x0 := 1 / float64(len(gs))
	r, err := Float64(g.intner)
	if err != nil {
		return 0, err
	}
	i := 0
	for i < len(gs) && gs[i] < r {
		i++
	}
	return (r-gs[i-1])/(gs[i]-gs[i-1])*x0 + float64(i-1)*x0, nil
}
