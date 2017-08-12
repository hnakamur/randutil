package randutil

import "math"

// Exp returns an exponential distributed random value.
func Exp(intner Intner, lambda float64) (float64, error) {
	t, err := Float64(intner)
	if err != nil {
		return 0, err
	}
	return -1 / lambda * math.Log(1-t), nil
}
