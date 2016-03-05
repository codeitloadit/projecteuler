package main

import (
	"fmt"
)

func main() {
	for c := 1; c <= 1000; c++ {
		for b := 1; b <= 1000; b++ {
			for a := 1; a <= 1000; a++ {
				if a*a+b*b == c*c && a+b+c == 1000 {
					fmt.Println(a * b * c)
					return
				}
			}
		}
	}
}
