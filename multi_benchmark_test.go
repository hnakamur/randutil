package randutil

import (
	"testing"

	"golang.org/x/exp/rand"
)

func BenchmarkMultiIntnNoDup(b *testing.B) {
	b.Run("m3n8", func(b *testing.B) {
		const m = 3
		const n = 8
		s := rand.NewSource(NewSeed())
		for i := 0; i < b.N; i++ {
			MultiIntnNoDup(s, m, n)
		}
	})
	b.Run("m3n100000", func(b *testing.B) {
		const m = 3
		const n = 100000
		s := rand.NewSource(NewSeed())
		for i := 0; i < b.N; i++ {
			MultiIntnNoDup(s, m, n)
		}
	})
	b.Run("m100n100000", func(b *testing.B) {
		const m = 100
		const n = 100000
		s := rand.NewSource(NewSeed())
		for i := 0; i < b.N; i++ {
			MultiIntnNoDup(s, m, n)
		}
	})
	b.Run("m1000n100000", func(b *testing.B) {
		const m = 1000
		const n = 100000
		s := rand.NewSource(NewSeed())
		for i := 0; i < b.N; i++ {
			MultiIntnNoDup(s, m, n)
		}
	})
	b.Run("m10000n100000", func(b *testing.B) {
		const m = 10000
		const n = 100000
		s := rand.NewSource(NewSeed())
		for i := 0; i < b.N; i++ {
			MultiIntnNoDup(s, m, n)
		}
	})
	b.Run("m100000n100000", func(b *testing.B) {
		const m = 100000
		const n = 100000
		s := rand.NewSource(NewSeed())
		for i := 0; i < b.N; i++ {
			MultiIntnNoDup(s, m, n)
		}
	})
}

func BenchmarkMultiIntnNoDupByRanges(b *testing.B) {
	b.Run("m3n8", func(b *testing.B) {
		const m = 3
		const n = 8
		s := rand.NewSource(NewSeed())
		for i := 0; i < b.N; i++ {
			multiIntnNoDupByRanges(s, m, n)
		}
	})
	b.Run("m3n100000", func(b *testing.B) {
		const m = 3
		const n = 100000
		s := rand.NewSource(NewSeed())
		for i := 0; i < b.N; i++ {
			multiIntnNoDupByRanges(s, m, n)
		}
	})
	b.Run("m100n100000", func(b *testing.B) {
		const m = 100
		const n = 100000
		s := rand.NewSource(NewSeed())
		for i := 0; i < b.N; i++ {
			multiIntnNoDupByRanges(s, m, n)
		}
	})
	b.Run("m1000n100000", func(b *testing.B) {
		const m = 1000
		const n = 100000
		s := rand.NewSource(NewSeed())
		for i := 0; i < b.N; i++ {
			multiIntnNoDupByRanges(s, m, n)
		}
	})
	b.Run("m10000n100000", func(b *testing.B) {
		const m = 10000
		const n = 100000
		s := rand.NewSource(NewSeed())
		for i := 0; i < b.N; i++ {
			multiIntnNoDupByRanges(s, m, n)
		}
	})
	b.Run("m100000n100000", func(b *testing.B) {
		const m = 100000
		const n = 100000
		s := rand.NewSource(NewSeed())
		for i := 0; i < b.N; i++ {
			multiIntnNoDupByRanges(s, m, n)
		}
	})
}

func BenchmarkMultiIntnNoDupByShuffle(b *testing.B) {
	b.Run("m3n8", func(b *testing.B) {
		const m = 3
		const n = 8
		s := rand.NewSource(NewSeed())
		for i := 0; i < b.N; i++ {
			multiIntnNoDupByShuffle(s, m, n)
		}
	})
	b.Run("m3n100000", func(b *testing.B) {
		const m = 3
		const n = 100000
		s := rand.NewSource(NewSeed())
		for i := 0; i < b.N; i++ {
			multiIntnNoDupByShuffle(s, m, n)
		}
	})
	b.Run("m100n100000", func(b *testing.B) {
		const m = 100
		const n = 100000
		s := rand.NewSource(NewSeed())
		for i := 0; i < b.N; i++ {
			multiIntnNoDupByShuffle(s, m, n)
		}
	})
	b.Run("m1000n100000", func(b *testing.B) {
		const m = 1000
		const n = 100000
		s := rand.NewSource(NewSeed())
		for i := 0; i < b.N; i++ {
			multiIntnNoDupByShuffle(s, m, n)
		}
	})
	b.Run("m10000n100000", func(b *testing.B) {
		const m = 10000
		const n = 100000
		s := rand.NewSource(NewSeed())
		for i := 0; i < b.N; i++ {
			multiIntnNoDupByShuffle(s, m, n)
		}
	})
	b.Run("m100000n100000", func(b *testing.B) {
		const m = 100000
		const n = 100000
		s := rand.NewSource(NewSeed())
		for i := 0; i < b.N; i++ {
			multiIntnNoDupByShuffle(s, m, n)
		}
	})
}

func multiIntnNoDupByShuffle(src rand.Source, m, n int) []int {
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
