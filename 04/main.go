package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(n int) bool {
	fwd := strconv.Itoa(n)
	rev := reverse(fwd)
	return fwd == rev
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	var highest int
	for i := 999; i > 0; i-- {
		for j := 999; j > 0; j-- {
			product := i * j
			if product > highest && isPalindrome(product) {
				highest = product
			}
		}
	}
	fmt.Println(highest)
}
