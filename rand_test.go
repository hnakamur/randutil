package randutil

import (
	"math/rand"
	"os"
	"strconv"
	"testing"
)

const (
	testCount = 100
)

func TestRandDigit95OneByOne(t *testing.T) {
	r := rand.New(rand.NewSource(seed))
	results := make([]int, batchSize)
	buf := make([]byte, 0, 80)
	for i := 0; i < testCount; i++ {
		buf = buf[:0]
		for j := 0; j < batchSize; j++ {
			results[j] = r.Intn(digitCount)

			if j > 0 {
				buf = append(buf, ' ')
			}
			if results[j] < 10 {
				buf = append(buf, '0')
			}
			buf = strconv.AppendUint(buf, uint64(results[j]), 10)
		}
		buf = append(buf, '\n')
		os.Stdout.Write(buf)
	}
}

func TestRandDigit95Batch(t *testing.T) {
	r := rand.New(rand.NewSource(seed))

	n := 1
	for j := 0; j < batchSize; j++ {
		n *= digitCount
	}

	results := make([]int, batchSize)
	buf := make([]byte, 0, 80)
	for i := 0; i < testCount; i++ {
		v := r.Intn(n)
		buf = buf[:0]
		for j := 0; j < batchSize; j++ {
			results[j] = v % digitCount
			v = v / digitCount

			if j > 0 {
				buf = append(buf, ' ')
			}
			if results[j] < 10 {
				buf = append(buf, '0')
			}
			buf = strconv.AppendUint(buf, uint64(results[j]), 10)
		}
		buf = append(buf, '\n')
		os.Stdout.Write(buf)
	}
}

func TestRandDigit95Batch2(t *testing.T) {
	r := rand.New(rand.NewSource(seed))

	n := 1
	for j := 0; j < batchSize; j++ {
		n *= digitCount
	}

	results := make([]int, batchSize)
	buf := make([]byte, 0, 80)
	for i := 0; i < testCount; i++ {
		v := r.Intn(n)
		buf = buf[:0]
		for j := 0; j < batchSize; j++ {
			v2 := v / digitCount
			results[j] = v - v2*digitCount
			v = v2

			if j > 0 {
				buf = append(buf, ' ')
			}
			if results[j] < 10 {
				buf = append(buf, '0')
			}
			buf = strconv.AppendUint(buf, uint64(results[j]), 10)
		}
		buf = append(buf, '\n')
		os.Stdout.Write(buf)
	}
}
