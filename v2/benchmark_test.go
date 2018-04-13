package v2_test

import (
	"math"
	"testing"

	randutil "github.com/hnakamur/randutil/v2"
)

func BenchmarkMathIntner(b *testing.B) {
	intner := randutil.NewMathIntner(0x1122334455667788)
	for i := 0; i < b.N; i++ {
		intner.Intn(1024)
	}
}

func BenchmarkXorShift64StarIntner(b *testing.B) {
	intner := randutil.NewXorShift64StarIntner(0x1122334455667788)
	for i := 0; i < b.N; i++ {
		intner.Intn(1024)
	}
}

func BenchmarkXorShift128PlusIntner(b *testing.B) {
	intner := randutil.NewXorShift128PlusIntner(0x1122334455667788)
	for i := 0; i < b.N; i++ {
		intner.Intn(1024)
	}
}

func BenchmarkXorShift1024StarIntner(b *testing.B) {
	intner := randutil.NewXorShift1024StarIntner(0x1122334455667788)
	for i := 0; i < b.N; i++ {
		intner.Intn(1024)
	}
}

func BenchmarkMathIntnerMaxInt64(b *testing.B) {
	intner := randutil.NewMathIntner(0x1122334455667788)
	for i := 0; i < b.N; i++ {
		intner.Intn(math.MaxInt64)
	}
}

func BenchmarkXorShift64StarIntnerMaxInt64(b *testing.B) {
	intner := randutil.NewXorShift64StarIntner(0x1122334455667788)
	for i := 0; i < b.N; i++ {
		intner.Intn(math.MaxInt64)
	}
}

func BenchmarkXorShift128PlusIntnerMaxInt64(b *testing.B) {
	intner := randutil.NewXorShift128PlusIntner(0x1122334455667788)
	for i := 0; i < b.N; i++ {
		intner.Intn(math.MaxInt64)
	}
}

func BenchmarkXorShift1024StarIntnerMaxInt64(b *testing.B) {
	intner := randutil.NewXorShift1024StarIntner(0x1122334455667788)
	for i := 0; i < b.N; i++ {
		intner.Intn(math.MaxInt64)
	}
}
