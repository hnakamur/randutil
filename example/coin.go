package main

import (
	"fmt"
	"time"

	"github.com/hnakamur/randutil"
)

func main() {
	//randutil.NewCryptoIntner(),
	c, err := randutil.NewChooser(
		randutil.NewMathIntner(time.Now().UnixNano()),
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
