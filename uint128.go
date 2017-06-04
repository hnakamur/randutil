package randutil

import "math/big"

const digits = "0123456789abcdefghijklmnopqrstuvwxyz"

func AppendUint128InZeroPaddedBase36(buf []byte, v [16]byte) []byte {
	var tmp [randomIDEncodedByteLen]byte
	p := new(big.Int).SetBytes(v[:]).Append(tmp[:0], base)
	for i := randomIDEncodedByteLen - len(p) - 1; i >= 0; i-- {
		buf = append(buf, '0')
	}
	return append(buf, p...)
}

func Uint128ToZeroPaddedBase36(v [16]byte) string {
	var buf [randomIDEncodedByteLen]byte
	return string(AppendUint128InZeroPaddedBase36(buf[:0], v))
}
