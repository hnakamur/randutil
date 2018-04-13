// Package randutil/v2 provides a utility function to produce
// random integer, float, bytes, and texts.
// This package is meant for generating random inputs to be
// used for tests or simulations.
// It does NOT provide crypto level randomness.
//
// The randutil/v2 drops Intner using crypto/rand and PCG.
// It also drops Chooser.
package v2
