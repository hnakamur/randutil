package randutil

import (
	"errors"
	"fmt"
)

// General returns a general discrete distributed random value.
func General(intner Intner, g []float64) (float64, error) {
	if len(g) < 2 {
		return 0, errors.New("length of g must be greater than or equal to 2")
	}
	if g[0] != 0 {
		return 0, errors.New("the first element of g must be 0")
	}
	if g[len(g)-1] != 1 {
		return 0, errors.New("the last element of g must be 1")
	}
	for i := 1; i < len(g)-1; i++ {
		if g[i] < g[i-1] {
			return 0, fmt.Errorf("g must be monotonously increasing, g[%d]=%f < g[%d]=%f", i, g[i], i-1, g[i-1])
		}
	}

	// http://www.ishikawa-lab.com/montecarlo/4shou.html
	x0 := 1 / float64(len(g))
	r, err := Float64(intner)
	if err != nil {
		return 0, err
	}
	i := 0
	for i < len(g) && g[i] < r {
		i++
	}
	return (r-g[i-1])/(g[i]-g[i-1])*x0 + float64(i-1)*x0, nil
}
