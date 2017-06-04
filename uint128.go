package randutil

import "encoding/binary"

func AppendUint128InZeroPaddedBase36(buf []byte, v [16]byte) []byte {
	var tmp [randomIDEncodedByteLen]byte
	i := len(tmp) - 1
	n := bigEndianUint128(v[:])
	for !n.IsZero() {
		var r uint8
		n, r = divmodUint128ByUint8(n, base)
		tmp[i] = digits[r]
		i--
	}
	for ; i >= 0; i-- {
		tmp[i] = '0'
	}
	return append(buf, tmp[:]...)
}

func Uint128ToZeroPaddedBase36(v [16]byte) string {
	var buf [randomIDEncodedByteLen]byte
	return string(AppendUint128InZeroPaddedBase36(buf[:0], v))
}

type uint128 struct {
	u [2]uint64
}

func bigEndianUint128(v []byte) uint128 {
	return uint128{
		[2]uint64{
			binary.BigEndian.Uint64(v[:8]),
			binary.BigEndian.Uint64(v[8:16]),
		},
	}
}

const digits = "0123456789abcdefghijklmnopqrstuvwxyz"

func (u *uint128) AppendHex(buf []byte) []byte {
	var tmp [32]byte
	i := len(tmp)
	if u.u[0] == 0 {
		n := u.u[1]
		if n == 0 {
			return append(buf, '0')
		}

		for n > 0 {
			i--
			tmp[i] = digits[n&0x0f]
			n >>= 4
		}
		return append(buf, tmp[i:]...)
	} else {
		n := u.u[1]
		for j := 15; j >= 0; j-- {
			i--
			tmp[i] = digits[n&0x0f]
			n >>= 4
		}

		n = u.u[0]
		for n > 0 {
			i--
			tmp[i] = digits[n&0x0f]
			n >>= 4
		}
		return append(buf, tmp[i:]...)
	}
}

func (u *uint128) Hex() string {
	var buf [32]byte
	return string(u.AppendHex(buf[:]))
}

// https://en.wikipedia.org/wiki/Division_algorithm#Integer_division_.28unsigned.29_with_remainder
func divmodUint128ByUint8(n uint128, d uint8) (q uint128, r uint8) {
	if d == 0 {
		panic("division by zero")
	}

	if n.u[0] != 0 {
		m := uint64(0x8000000000000000)
		for i := 63; i >= 0; i-- {
			r <<= 1
			if n.u[0]&m != 0 {
				r |= 1
			}
			if r >= d {
				r -= d
				q.u[0] |= m
			}
			m >>= 1
		}
	}
	m := uint64(0x8000000000000000)
	for i := 63; i >= 0; i-- {
		r <<= 1
		if n.u[1]&m != 0 {
			r |= 1
		}
		if r >= d {
			r -= d
			q.u[1] |= m
		}
		m >>= 1
	}
	return
}

func (n *uint128) IsZero() bool {
	return n.u[0] == 0 && n.u[1] == 0
}
