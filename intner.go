package randutil

// Intner is the interface that returns a non-negative pseudo random int in the specified range.
//
// Intn returns, as an int, a non-negative pseudo-random number in [0,n). It panics if n <= 0.
//
// Uint64 returns a pseudo-random 64-bit value as a uint64.
type Intner interface {
	Intn(n int) (int, error)
	Uint64() (uint64, error)
}
