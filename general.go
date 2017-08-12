package randutil

// General returns a general discrete distributed random value.
func General(intner Intner, g []float64) (float64, error) {
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
