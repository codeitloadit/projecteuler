// Prime factorization derived from slightly modified version
// of sieve.go in Go source distribution.
package main

import (
	"fmt"
)

// Generate numbers until the limit max.
// after the 2, all the prime numbers are odd
// Send a channel signal when the limit is reached
func Generate(max int, ch chan<- int) {
	ch <- 2
	for i := 3; i <= max; i += 2 {
		ch <- i
	}
	ch <- -1 // signal that the limit is reached
}

// Filter copies the values from channel 'in' to channel 'out',
// removing those divisible by 'prime'.
func Filter(in <-chan int, out chan<- int, prime int) {
	for i := <-in; i != -1; i = <-in {
		if i%prime != 0 {
			out <- i
		}
	}
	out <- -1
}

// CalcPrimeFactors calculates the prime factors of an int.
func CalcPrimeFactors(n int) []int {
	rv := []int{}
	ch := make(chan int)
	go Generate(n, ch)
	for prime := <-ch; (prime != -1) && (n > 1); prime = <-ch {
		for n%prime == 0 {
			n = n / prime
			rv = append(rv, prime)
		}
		ch1 := make(chan int)
		go Filter(ch, ch1, prime)
		ch = ch1
	}
	return rv
}

func main() {
	factors := CalcPrimeFactors(600851475143)
	fmt.Println(factors[len(factors)-1])
}
