package main

import (
	"fmt"
)

//probleem descriptie :
// "You are climbing a staircase. It takes n steps to reach the top.
// Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?"

// Voorbeeld :
// Als input :  n = 3
// uitkomst : 3
// uitgetekend :
// 1. 1 step + 1 step + 1 step
// 2. 1 step + 2 steps
// 3. 2 steps + 1 step
// er zijn dus 3 manieren om met 1 of 2 stap naar n=3 te komen

func main() {

	var n uint64 = 5

	fmt.Println(ClimbingStairs(n))
}

func ClimbingStairs(n uint64) uint64 {
	var one, two uint64 = 1, 0

	for i := n; i > 0; i-- {
		temp := one
		one = one + two
		two = temp
	}

	return one
}
