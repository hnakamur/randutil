package randutil

type Intner interface {
	Intn(n int) (int, error)
}
