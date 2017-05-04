package randutil_test

import (
	"fmt"
	"time"

	"github.com/hnakamur/randutil"
)

func ExampleNewMathIntner() {
	intner := randutil.NewMathIntner(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		v, err := intner.Intn(6)
		if err != nil {
			panic(err)
		}
		fmt.Printf("dice was %d\n", v+1)
	}
}

func ExampleNewCryptoIntner() {
	intner := randutil.NewCryptoIntner()
	for i := 0; i < 10; i++ {
		v, err := intner.Intn(6)
		if err != nil {
			panic(err)
		}
		fmt.Printf("dice was %d\n", v+1)
	}
}

func ExampleChooser_Choose() {
	c, err := randutil.NewChooser(
		randutil.NewCryptoIntner(),
		[]randutil.Choice{
			{Weight: 1, Item: "head"},
			{Weight: 2, Item: "tail"},
		})
	if err != nil {
		panic(err)
	}

	counts := map[string]int{
		"head": 0,
		"tail": 0,
	}
	for i := 0; i < 10000; i++ {
		v, err := c.Choose()
		if err != nil {
			panic(err)
		}
		counts[v.(string)]++
	}
	fmt.Printf("head=%d, tail=%d\n", counts["head"], counts["tail"])
}
