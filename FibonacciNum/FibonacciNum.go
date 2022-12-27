package main

import (
	"fmt"
)

func main() {

	var n int = 20

	
	// var nums = []int{2, 7, 11, 15}

	// fmt.Println(HBOguy(nums))
	fmt.Println(fib(n))

}
// func HBOguy(nums []int) int  {


// 	totaal := 0

// 	for i := 0; i < len(nums); i++ {
// 		totaal += nums[i]
// 	}
	
// 	return totaal/len(nums)
// }


func fib(n int) int {
	zero, one := 0, 1

	for i := n; i > 0; i-- {
		temp := zero


		zero = zero + one
		one = temp
		// fmt.Println("second", temp, one, "res", zero)

	}

	return zero
}

