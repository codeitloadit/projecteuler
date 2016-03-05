package main

import (
	"fmt"
)

func main() {
	n := 20
	for {
		valid := true
		for i := 2; i <= 20; i++ {
			if n%i != 0 {
				valid = false
				break
			}
		}
		if valid {
			fmt.Println(n)
			return
		}
		n++
	}
}
