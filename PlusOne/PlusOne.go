package main

import (
	"fmt"
)

func main() {

	var digits = []int{9, 9, 9}

	fmt.Println(plusOne(digits))

}

func plusOne(digits []int) []int {
	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		digits[i], digits[j] = digits[j], digits[i]
	}
	// de for loop hierboven is de enige manier om een array te reversen lool

	one, i := 1, 0
	for one == 1 {
		if i < len(digits) {
			if digits[i] == 9 {
				digits[i] = 0
			} else {
				digits[i] += 1
				one = 0
			}
		} else {
			digits = append(digits, 1)
			one = 0
		}
		i++
	}
	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		digits[i], digits[j] = digits[j], digits[i]
	}
	return digits

}
