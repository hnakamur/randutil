package randutil

import (
	"sort"

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
	ret := make([]int, 0, m)
	for i := 0; i < m; i++ {
		r := rnd.Intn(n)
		v := a[r]
		ret = append(ret, v)

		// a = append(a[:r], a[r+1:]...)
		a = a[:r+copy(a[r:], a[r+1:])]
		n--
	}
	return ret
}

// multiIntnNoDup returns m multiple random integers
// of range [0, n) without duplication.
func multiIntnNoDupByRanges(src rand.Source, m, n int) []int {
	if src == nil {
		panic("src must not be nil")
	}
	if m > n {
		panic("m must be less than or equal to n")
	}
	rnd := rand.New(src)
	ret := make([]int, 0, m)
	ranges := intRanges{{start: 0, stop: n}}
	for i := 0; i < m; i++ {
		r := rnd.Intn(n)
		v := ranges.IndexToValue(r)
		ret = append(ret, v)

		ranges.RemoveInt(v)
		n--
	}
	return ret
}

type intRange struct {
	// start of range, inclusive
	start int
	// end of range, exclusive
	stop int
}

// intRanges is a multiple of ranges in ascending order.
type intRanges []intRange

func (rr *intRanges) RemoveInt(x int) {
	i := sort.Search(len(*rr), func(i int) bool {
		return (*rr)[i].start >= x
	})
	if i < len(*rr) {
		r := (*rr)[i]
		if x == r.start {
			if r.start+1 == r.stop {
				// delete (*rr)[i]
				*rr = append((*rr)[:i], (*rr)[i+1:]...)
				return
			}

			(*rr)[i].start++
			return
		}
	}

	i--
	if i < 0 {
		return
	}
	r := (*rr)[i]
	if r.stop <= x {
		return
	}

	if x == r.stop-1 {
		(*rr)[i].stop--
		return
	}

	// insert intRange{start: x+1, stop, r.stop} at i+1
	*rr = append(*rr, intRange{})
	copy((*rr)[(i+1)+1:], (*rr)[i+1:])
	(*rr)[i+1] = intRange{start: x + 1, stop: r.stop}

	(*rr)[i].stop = x
}

func (rr *intRanges) IndexToValue(i int) int {
	for _, r := range *rr {
		if i < r.stop-r.start {
			return r.start + i
		}
		i -= r.stop - r.start
	}
	panic("index out of bounds")
}
