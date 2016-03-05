package main

import (
	"fmt"
)

func main() {
	sum := 2
	terms := []int{1, 2}
	for {
		newTerm := terms[len(terms)-2] + terms[len(terms)-1]
		if newTerm > 4000000 {
			break
		}
		terms = append(terms, newTerm)
		if newTerm%2 == 0 {
			sum += newTerm
		}
	}
	fmt.Println(sum)
}
