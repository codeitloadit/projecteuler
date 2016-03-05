package main

import (
	"fmt"
)

var x = 100

func sumOfSquares(n int) int {
	var r int
	for i := 1; i <= x; i++ {
		r += i * i
	}
	return r
}

func squareOfSums(n int) int {
	var s int
	for i := 1; i <= x; i++ {
		s += i
	}
	return s * s
}

func main() {
	fmt.Println(squareOfSums(x) - sumOfSquares(x))
}
