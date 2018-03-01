package randutil

import (
	"math"
	"math/rand"
	"testing"
	"time"
)

const (
	digitCount = 95
)

var seed = time.Now().UnixNano()
var batchSize = int(math.Log2(math.MaxUint64) / math.Log2(digitCount))

func BenchmarkRandDigit95OneByOne(b *testing.B) {
	r := rand.New(rand.NewSource(seed))
	results := make([]int, batchSize)
	for i := 0; i < b.N; i++ {
		for j := 0; j < batchSize; j++ {
			results[j] = r.Intn(digitCount)
		}
	}
}

func BenchmarkRandDigit95Batch(b *testing.B) {
	r := rand.New(rand.NewSource(seed))

	n := 1
	for j := 0; j < batchSize; j++ {
		n *= digitCount
	}

	results := make([]int, batchSize)
	for i := 0; i < b.N; i++ {
		v := r.Intn(n)
		for j := 0; j < batchSize; j++ {
			results[j] = v % digitCount
			v = v / digitCount
		}
	}
}

func BenchmarkRandDigit95Batch2(b *testing.B) {
	r := rand.New(rand.NewSource(seed))

	n := 1
	for j := 0; j < batchSize; j++ {
		n *= digitCount
	}

	results := make([]int, batchSize)
	for i := 0; i < b.N; i++ {
		v := r.Intn(n)
		for j := 0; j < batchSize; j++ {
			v2 := v / digitCount
			results[j] = v - v2*digitCount
			v = v2
		}
	}
}
