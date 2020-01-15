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

func TestMultiIntnNoDupByShuffle(t *testing.T) {
	const tries = 100_000

	// This test targets the case m == n
	const n = 8
	want := make([]int, n)
	for i := 0; i < n; i++ {
		want[i] = i
	}

	s := rand.NewSource(NewSeed())
	for i := 0; i < tries; i++ {
		values := multiIntnNoDupByShuffle(s, n, n)
		got := append([]int{}, values...)
		sort.Ints(got)
		if !reflect.DeepEqual(got, want) || got[0] != 0 || got[n-1] != n-1 {
			t.Errorf("unexpected result, got:%v", got)
		}
	}
}

func TestIntRanges_RemoveInt(t *testing.T) {
	testCases := []struct {
		ranges intRanges
		x      int
		want   intRanges
	}{
		{ranges: intRanges{{start: 0, stop: 1}}, x: 0, want: intRanges{}},
		{ranges: intRanges{{start: 0, stop: 2}}, x: 0, want: intRanges{{start: 1, stop: 2}}},
		{ranges: intRanges{{start: 0, stop: 2}}, x: 1, want: intRanges{{start: 0, stop: 1}}},
		{ranges: intRanges{{start: 0, stop: 3}},
			x:    1,
			want: intRanges{{start: 0, stop: 1}, {start: 2, stop: 3}}},
		{ranges: intRanges{{start: 0, stop: 1}, {start: 2, stop: 3}},
			x:    0,
			want: intRanges{{start: 2, stop: 3}}},
		{ranges: intRanges{{start: 0, stop: 1}, {start: 2, stop: 3}},
			x:    2,
			want: intRanges{{start: 0, stop: 1}}},
		{ranges: intRanges{{start: 0, stop: 1}, {start: 2, stop: 4}},
			x:    2,
			want: intRanges{{start: 0, stop: 1}, {start: 3, stop: 4}}},
		{ranges: intRanges{{start: 0, stop: 1}, {start: 2, stop: 4}},
			x:    3,
			want: intRanges{{start: 0, stop: 1}, {start: 2, stop: 3}}},
		{ranges: intRanges{{start: 0, stop: 1}, {start: 2, stop: 5}},
			x:    3,
			want: intRanges{{start: 0, stop: 1}, {start: 2, stop: 3}, {start: 4, stop: 5}}},
		{ranges: intRanges{{start: 0, stop: 1}, {start: 2, stop: 5}, {start: 6, stop: 7}},
			x:    3,
			want: intRanges{{start: 0, stop: 1}, {start: 2, stop: 3}, {start: 4, stop: 5}, {start: 6, stop: 7}}},
	}
	for i, c := range testCases {
		c.ranges.RemoveInt(c.x)
		if !reflect.DeepEqual(c.ranges, c.want) {
			t.Errorf("unmatch, caseID=%d, got=%v, want=%v", i, c.ranges, c.want)
		}
	}
}

func TestIntRanges_IndexToValue(t *testing.T) {
	testCases := []struct {
		ranges intRanges
		i      int
		want   int
	}{
		{ranges: intRanges{{start: 0, stop: 1}}, i: 0, want: 0},
		{ranges: intRanges{{start: 0, stop: 2}}, i: 1, want: 1},
		{ranges: intRanges{{start: 0, stop: 1}, {start: 2, stop: 4}}, i: 0, want: 0},
		{ranges: intRanges{{start: 0, stop: 1}, {start: 2, stop: 4}}, i: 1, want: 2},
		{ranges: intRanges{{start: 0, stop: 1}, {start: 2, stop: 4}}, i: 2, want: 3},
	}
	for i, c := range testCases {
		got := c.ranges.IndexToValue(c.i)
		if got != c.want {
			t.Errorf("unmatch, caseID=%d, got=%v, want=%v", i, got, c.want)
		}
	}
}
