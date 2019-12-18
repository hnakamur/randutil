package randutil

import (
	"sort"

	"golang.org/x/exp/rand"
)

// Weighted is a weighted random chooser.
type Weighted struct {
	rnd     *rand.Rand
	weights []float64
}

// NewWeighted creates a weighted random chooser.
// If src is nil, the rand.Float will be used.
func NewWeighted(src rand.Source, weights []float64) *Weighted {
	if src == nil {
		panic("src must not be nil")
	}
	if len(weights) == 0 {
		panic("weights must not be empty")
	}

	w := &Weighted{
		rnd:     rand.New(src),
		weights: make([]float64, len(weights)),
	}

	var total float64
	for i, weight := range weights {
		if weight < 0 {
			panic("weight must be non-negative")
		}
		total += weight
		w.weights[i] = total
	}
	for i := range w.weights {
		w.weights[i] /= total
	}

	return w
}

// Choose chooses one value from values.
func (w *Weighted) Choose(values []string) string {
	if len(values) != len(w.weights) {
		panic("values length must be equals to weights")
	}

	x := w.rnd.Float64()
	i := sort.Search(len(w.weights), func(i int) bool {
		return w.weights[i] >= x
	})
	return values[i]
}
