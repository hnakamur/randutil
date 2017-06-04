package randutil_test

import (
	"testing"

	"github.com/hnakamur/randutil"
)

func BenchmarkRandomID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := randutil.RandomID(nil)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkUint128ToZeroPaddedBase36(b *testing.B) {
	values := [][16]byte{
		{'\xfb', '\x58', '\x55', '\x48', '\xec', '\x72', '\x9d', '\x11', '\xe7', '\x3b', '\x68', '\xe4', '\x4f', '\x62', '\x7c', '\xb9'},
		{'\x0a', '\xe7', '\xc9', '\x8b', '\xdb', '\x66', '\xf8', '\x6d', '\x23', '\x9c', '\x1d', '\x1c', '\x85', '\x2b', '\x03', '\xcd'},
		{'\x5b', '\xe2', '\xdf', '\xef', '\x39', '\xba', '\x59', '\x0f', '\xd0', '\xef', '\x1c', '\x75', '\x95', '\x02', '\x8b', '\x52'},
		{'\xe3', '\x8b', '\x61', '\xe0', '\x3e', '\x26', '\x16', '\xa8', '\x7f', '\x17', '\xc8', '\xc6', '\x48', '\xac', '\x86', '\xe1'},
		{'\xa4', '\x40', '\x67', '\xc7', '\x61', '\x2c', '\x67', '\x92', '\xb5', '\x42', '\x08', '\x15', '\x48', '\x26', '\x06', '\xf1'},
		{'\x60', '\x07', '\xf9', '\x83', '\x13', '\xdb', '\xdc', '\x38', '\x14', '\x50', '\xed', '\x73', '\xeb', '\xf4', '\x23', '\x38'},
		{'\xc5', '\x61', '\x0a', '\x7d', '\x0b', '\x11', '\x87', '\x7b', '\x96', '\x9d', '\xd1', '\xf4', '\xfb', '\x1d', '\x0c', '\x95'},
		{'\xcf', '\x72', '\x3e', '\x3c', '\xd2', '\x16', '\xcb', '\x7c', '\xdc', '\x32', '\x77', '\xa3', '\xd2', '\xab', '\x49', '\xbc'},
		{'\xac', '\x98', '\x8c', '\x57', '\x30', '\xf2', '\xce', '\x25', '\x5c', '\xad', '\x69', '\x5e', '\xf0', '\x0b', '\xa3', '\x05'},
		{'\x79', '\x79', '\x7b', '\xe1', '\x1d', '\x58', '\x09', '\x56', '\xc1', '\xab', '\xf2', '\xbf', '\x12', '\x4a', '\x30', '\x3c'},
	}
	for i := 0; i < b.N; i++ {
		for _, v := range values {
			randutil.Uint128ToZeroPaddedBase36(v)
		}
	}
}
