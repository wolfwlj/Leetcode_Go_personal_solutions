package main

import (
	"fmt"
)


func main() {
	
	prices := []int{3,2,6,5,0,3}

	// prices := []int{7,6,4,3,1}
	// prices := []int{2,1,2,0,1}
	// prices := []int{2,4,1}
	
	// s := "string"
	fmt.Println(maxProfit(prices))

}
func maxProfit(prices []int) int {
	// highestProfit := 0
	// lowest := 0
	// highest := 0
	// // len := len(prices)
	// for i := 0; i < len(prices); i++ {

	// 	if prices[i] > highest {
	// 		highest = prices[i]
	// 		fmt.Println(i)

	// 	} 
	// 	if i == 0 {lowest = prices[i]}
	// 	if prices[i] < lowest  {

	// 		if i+1 != len(prices) {
	// 			lowest = prices[i]
	// 			highest = 0
				
	// 		}

	// 	}

	// 	// A comment.
	// 	// fmt.Println(i+1, len(prices))
	// 	fmt.Println(i, highest, lowest)

	// 	if (highest - lowest > highestProfit) {
	// 		highestProfit = highest - lowest
	// 	}


	// }
	
	// if highestProfit > 0 {return highestProfit}

	// return 0

	// better solution, less run time and memory
	l := 0
	r := 1
	highestProfit := 0

	for r < len(prices) {
		if prices[l] < prices[r] {
			profit := prices[r] - prices[l]
			highestProfit = Max(highestProfit, profit)
		} else {
			l = r
		}
		r +=1
	}
	return highestProfit
}


func Max(x, y int) int {
    if x < y {
        return y
    }
    return x
}