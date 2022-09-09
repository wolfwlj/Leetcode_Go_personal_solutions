package main

import (
	"fmt"
)

func main() {

	var n int = 20

	fmt.Println(fib(n))
}

func fib(n int) int {
	zero, one := 0, 1

	for i := n; i > 0; i-- {
		temp := zero

		fmt.Println("first", temp, one)

		zero = zero + one
		one = temp
		// fmt.Println("second", temp, one, "res", zero)
		fmt.Println("first", zero)

	}

	return zero
}
