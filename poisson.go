package randutil

import "math"

// Poisson returns a poisson distributed random value.
func Poisson(intner Intner, lambda float64) (int, error) {
	// http://www.ishikawa-lab.com/montecarlo/4shou.html
	k := 0
	xp, err := Float64(intner)
	if err != nil {
		return 0, err
	}
	el := math.Exp(-lambda)
	for xp >= el {
		x, err := Float64(intner)
		if err != nil {
			return 0, err
		}
		xp *= x
		k++
	}
	return k, nil
}
