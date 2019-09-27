[DEPRECATED] Use [golang.org/x/exp/rand#PCGSource](https://godoc.org/golang.org/x/exp/rand#PCGSource) and [gonum.org/v1/gonum/stat/distuv](https://godoc.org/gonum.org/v1/gonum/stat/distuv) instead.

randutil  [![GoDoc](https://godoc.org/github.com/hnakamur/randutil?status.png)](https://godoc.org/github.com/hnakamur/randutil)
========

This is a Go package for providing some utilities for random.

This Package randutil provides a utility function to get non-negative
pesudo-random int using math/rand or crypto/rand.

This package also provides Chooser which chooses an item from choices.
Each choice has a weight which affects possibility for it to be choosed.

## License

MIT
