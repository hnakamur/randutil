package randutil_test

import (
	"testing"

	"github.com/hnakamur/randutil"
)

func BenchmarkRandomID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := randutil.RandomID()
		if err != nil {
			b.Fatal(err)
		}
	}
}
