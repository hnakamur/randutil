package randutil

import (
	"testing"
)

func TestUint128AppendHex(t *testing.T) {
	testCases := []struct {
		v    [16]byte
		want string
	}{
		{
			[16]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			"0",
		},
		{
			[16]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
			"1",
		},
		{
			[16]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x0f},
			"f",
		},
		{
			[16]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x10},
			"10",
		},
		{
			[16]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x12},
			"12",
		},
		{
			[16]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0xff},
			"ff",
		},
		{
			[16]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0},
			"100",
		},
		{
			[16]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1},
			"101",
		},
		{
			[16]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0x10},
			"110",
		},
		{
			[16]byte{0, 0, 0, 0, 0, 0, 0, 0, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
			"ffffffffffffffff",
		},
		{
			[16]byte{0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
			"10000000000000000",
		},
		{
			[16]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xfe},
			"fffffffffffffffffffffffffffffffe",
		},
		{
			[16]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
			"ffffffffffffffffffffffffffffffff",
		},
	}
	for _, tc := range testCases {
		n := bigEndianUint128(tc.v[:])
		got := string(n.AppendHex(nil))
		if got != tc.want {
			t.Errorf("got %q, want %q", got, tc.want)
		}
	}
}

func TestUint128ToZeroPaddedBase36(t *testing.T) {
	testCases := []struct {
		v    [16]byte
		want string
	}{
		{
			[16]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			"0000000000000000000000000",
		},
		{
			[16]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
			"0000000000000000000000001",
		},
		{
			[16]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 35},
			"000000000000000000000000z",
		},
		{
			[16]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 36},
			"0000000000000000000000010",
		},
		{
			[16]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 5, 15},
			"00000000000000000000000zz",
		},
		{
			[16]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 5, 16},
			"0000000000000000000000100",
		},
		{
			[16]byte{0, 0, 0, 0, 0, 0, 0, 0, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
			"0000000000003w5e11264sgsf",
		},
		{
			[16]byte{0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
			"0000000000003w5e11264sgsg",
		},
		{
			[16]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xfe},
			"f5lxx1zz5pnorynqglhzmsp32",
		},
		{
			[16]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
			"f5lxx1zz5pnorynqglhzmsp33",
		},
	}
	for _, tc := range testCases {
		got := Uint128ToZeroPaddedBase36(tc.v)
		if got != tc.want {
			t.Errorf("got %s, want %s", got, tc.want)
		}
	}
}
