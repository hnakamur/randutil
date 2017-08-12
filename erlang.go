package randutil

import "math"

// Erlang returns an erlang distributed random value.
func Erlang(intner Intner, lambda float64, k int) (float64, error) {
	// http://www.ishikawa-lab.com/montecarlo/4shou.html
	tp := 1.0
	for n := 1; n <= k; n++ {
		f, err := Float64(intner)
		if err != nil {
			return 0, err
		}
		tp *= (1 - f)
	}
	return -1 / lambda / float64(k) * math.Log(tp), nil
}
