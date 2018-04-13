package v2

// Float64 returns, as a float64, a pseudo-random number in [0.0,1.0).
func Float64(intner Intner) float64 {
	// https://github.com/golang/go/blob/c007ce824d9a4fccb148f9204e04c23ed2984b71/src/math/rand/rand.go#L105-L106
	u := intner.Int63n(1 << 53)
	return float64(u) / (1 << 53)
}
