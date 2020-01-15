package randutil

import (
	"golang.org/x/exp/rand"
)

// MultiIntnNoDup returns m multiple random integers
// of range [0, n) without duplication.
func MultiIntnNoDup(src rand.Source, m, n int) []int {
	if src == nil {
		panic("src must not be nil")
	}
	if m > n {
		panic("m must be less than or equal to n")
	}

	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = i
	}

	rnd := rand.New(src)
	rnd.Shuffle(n, func(i, j int) { a[i], a[j] = a[j], a[i] })
	return append([]int{}, a[:m]...)
}
