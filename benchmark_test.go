package randutil

import (
	"math"
	"math/rand"
	"testing"
	"time"

	erand "golang.org/x/exp/rand"
)

func BenchmarkInt63n(b *testing.B) {
	seed := time.Now().UnixNano()

	b.Run("mathMaxInt64", func(b *testing.B) {
		const n = math.MaxInt64
		rnd := rand.New(rand.NewSource(seed))
		for i := 0; i < b.N; i++ {
			rnd.Int63n(n)
		}
	})
	b.Run("expMaxInt64", func(b *testing.B) {
		const n = math.MaxInt64
		rnd := erand.New(erand.NewSource(uint64(seed)))
		for i := 0; i < b.N; i++ {
			rnd.Int63n(n)
		}
	})
	b.Run("math2147483647", func(b *testing.B) {
		const n = 1<<31 - 1
		rnd := rand.New(rand.NewSource(seed))
		for i := 0; i < b.N; i++ {
			rnd.Int63n(n)
		}
	})
	b.Run("exp2147483647", func(b *testing.B) {
		const n = 1<<31 - 1
		rnd := erand.New(erand.NewSource(uint64(seed)))
		for i := 0; i < b.N; i++ {
			rnd.Int63n(n)
		}
	})
}

func BenchmarkInt31n(b *testing.B) {
	seed := time.Now().UnixNano()

	b.Run("mathMaxInt32", func(b *testing.B) {
		const n = math.MaxInt32
		rnd := rand.New(rand.NewSource(seed))
		for i := 0; i < b.N; i++ {
			rnd.Int31n(n)
		}
	})
	b.Run("expMaxInt32", func(b *testing.B) {
		const n = math.MaxInt32
		rnd := erand.New(erand.NewSource(uint64(seed)))
		for i := 0; i < b.N; i++ {
			rnd.Int31n(n)
		}
	})
	b.Run("math2147483646", func(b *testing.B) {
		const n = 1<<31 - 1 - 1
		rnd := rand.New(rand.NewSource(seed))
		for i := 0; i < b.N; i++ {
			rnd.Int31n(n)
		}
	})
	b.Run("exp2147483646", func(b *testing.B) {
		const n = 1<<31 - 1 - 1
		rnd := erand.New(erand.NewSource(uint64(seed)))
		for i := 0; i < b.N; i++ {
			rnd.Int31n(n)
		}
	})
}

func BenchmarkIntn(b *testing.B) {
	seed := time.Now().UnixNano()

	b.Run("mathMaxInt64", func(b *testing.B) {
		const n = math.MaxInt64
		rnd := rand.New(rand.NewSource(seed))
		for i := 0; i < b.N; i++ {
			rnd.Intn(n)
		}
	})
	b.Run("expMaxInt64", func(b *testing.B) {
		const n = math.MaxInt64
		rnd := erand.New(erand.NewSource(uint64(seed)))
		for i := 0; i < b.N; i++ {
			rnd.Intn(n)
		}
	})
	b.Run("math2147483647", func(b *testing.B) {
		const n = 1<<31 - 1
		rnd := rand.New(rand.NewSource(seed))
		for i := 0; i < b.N; i++ {
			rnd.Intn(n)
		}
	})
	b.Run("exp2147483647", func(b *testing.B) {
		const n = 1<<31 - 1
		rnd := erand.New(erand.NewSource(uint64(seed)))
		for i := 0; i < b.N; i++ {
			rnd.Intn(n)
		}
	})
	b.Run("mathMaxInt32", func(b *testing.B) {
		const n = math.MaxInt32
		rnd := rand.New(rand.NewSource(seed))
		for i := 0; i < b.N; i++ {
			rnd.Intn(n)
		}
	})
	b.Run("expMaxInt32", func(b *testing.B) {
		const n = math.MaxInt32
		rnd := erand.New(erand.NewSource(uint64(seed)))
		for i := 0; i < b.N; i++ {
			rnd.Intn(n)
		}
	})
	b.Run("math2147483646", func(b *testing.B) {
		const n = 1<<31 - 1 - 1
		rnd := rand.New(rand.NewSource(seed))
		for i := 0; i < b.N; i++ {
			rnd.Intn(n)
		}
	})
	b.Run("exp2147483646", func(b *testing.B) {
		const n = 1<<31 - 1 - 1
		rnd := erand.New(erand.NewSource(uint64(seed)))
		for i := 0; i < b.N; i++ {
			rnd.Intn(n)
		}
	})
}
