package randutil_test

import (
	"strings"
	"testing"
	"time"

	"github.com/hnakamur/randutil"
)

func TestEncoder_RandomString(t *testing.T) {
	intner := randutil.NewMathIntner(time.Now().Unix())
	digits := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	enc, err := randutil.NewEncoder(intner, []byte(digits))
	if err != nil {
		t.Fatal(err)
	}

	const length = 64
	const count = 100
	for i := uint(0); i < count; i++ {
		s, err := enc.RandomString(length)
		if err != nil {
			t.Fatal(err)
		}
		if len(s) != length {
			t.Errorf("unexpected length, got=%d, want=%d", len(s), length)
		}
		for _, r := range s {
			if !strings.ContainsRune(digits, r) {
				t.Errorf("unexpected rune in generated string, rune=%v", r)
			}
		}
	}
}

func BenchmarkEncoder_RandomString(b *testing.B) {
	intner := randutil.NewMathIntner(time.Now().Unix())
	digits := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	enc, err := randutil.NewEncoder(intner, []byte(digits))
	if err != nil {
		b.Fatal(err)
	}

	const length = 64
	for i := 0; i < b.N; i++ {
		_, err := enc.RandomString(length)
		if err != nil {
			b.Fatal(err)
		}
	}
}
