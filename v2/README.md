randutil/v2  [![GoDoc](https://godoc.org/github.com/hnakamur/randutil?status.png)](https://godoc.org/github.com/hnakamur/randutil/v2)
===============

This is a Go package for providing some utilities for random.

Package randutil/v2 provides a utility function to produce
random integer, float, bytes, and texts.
This package is meant for generating random inputs to be
used for tests or simulations.
It does NOT provide crypto level randomness.

The randutil/v2 drops Intner using crypto/rand and PCG.
It also drops Chooser.

## License

MIT
