package main

import (
	"fmt"

	"github.com/hnakamur/randutil"
)

func main() {
	//intner := randutil.NewMathIntner(3)
	intner := randutil.NewCryptoIntner()
	for i := 0; i < 10; i++ {
		v, err := intner.Intn(6)
		if err != nil {
			panic(err)
		}
		fmt.Printf("dice was %d\n", v+1)
	}
}
