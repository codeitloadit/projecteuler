package main

import (
	"fmt"
	"math/big"
)

func main() {
	var p int
	i := 1
	for {
		if big.NewInt(int64(i)).ProbablyPrime(2) {
			p++
		}
		if p == 10001 {
			fmt.Println(i)
			return
		}
		i++
	}
}
