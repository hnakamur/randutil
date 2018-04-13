package v2

// Intner is the interface that returns a non-negative pseudo random int in the specified range.
//
// Intn returns, as an int, a non-negative pseudo-random number in [0,n). It panics if n <= 0.
//
// Int63n returns, as an int64, a non-negative pseudo-random number in [0,n) from the default Source. It panics if n <= 0.
type Intner interface {
	Intn(n int) int
	Int63n(n int64) int64
}
