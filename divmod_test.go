package randutil

import "testing"

func TestDivmodUint64(t *testing.T) {
	testCases := []struct {
		n, d, q, r uint64
	}{
		{1, 1, 1, 0},
		{3, 1, 3, 0},
		{2, 2, 1, 0},
		{3, 2, 1, 1},
		{4, 2, 2, 0},
		{5, 2, 2, 1},
		{35, 36, 0, 35},
		{36, 36, 1, 0},
		{37, 36, 1, 1},
		{71, 36, 1, 35},
		{110, 36, 3, 2},
		{397, 36, 11, 1},
	}
	for _, tc := range testCases {
		q, r := divmodUint64(tc.n, tc.d)
		if q != tc.q || r != tc.r {
			t.Errorf("n=%d, d=%d, gotQ=%d, gotR=%d, wantQ=%d, wantR=%d",
				tc.n, tc.d, q, r, tc.q, tc.r)
		}
	}
}
