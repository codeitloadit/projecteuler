package main

import (
	"fmt"
	"math/big"
)

var d = 2000000
var sum int

func main() {
	for i := d; i > 1; i-- {
		if big.NewInt(int64(i)).ProbablyPrime(4) {
			sum += i
		}
	}
	fmt.Println(sum)
}
