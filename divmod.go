package randutil

// https://en.wikipedia.org/wiki/Division_algorithm#Integer_division_.28unsigned.29_with_remainder
func divmodUint64(n, d uint64) (q, r uint64) {
	if d == 0 {
		panic("division by zero")
	}
	m := uint64(0x8000000000000000)
	for i := 63; i >= 0; i-- {
		r <<= 1
		if n&m != 0 {
			r |= 1
		}
		if r >= d {
			r -= d
			q |= m
		}
		m >>= 1
	}
	return
}
