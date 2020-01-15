package randutil

import (
	"reflect"
	"sort"
	"testing"

	"golang.org/x/exp/rand"
)

func TestMultiIntnNoDup(t *testing.T) {
	const tries = 100_000

	// This test targets the case m == n
	const n = 8
	want := make([]int, n)
	for i := 0; i < n; i++ {
		want[i] = i
	}

	s := rand.NewSource(NewSeed())
	for i := 0; i < tries; i++ {
		values := MultiIntnNoDup(s, n, n)
		got := append([]int{}, values...)
		sort.Ints(got)
		if !reflect.DeepEqual(got, want) || got[0] != 0 || got[n-1] != n-1 {
			t.Errorf("unexpected result, got:%v", got)
		}
	}
}
