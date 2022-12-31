package main

import (
	"fmt"
)

func main() {
	// problem : 
	//Given an unsorted array of integers nums, return the length of the longest consecutive elements sequence.

	//You must write an algorithm that runs in O(n) time.




	var nums = []int{9,1,4,7,3,-1,0,5,8,-1,6}
	//9,1,4,7,3,-1,0,5,8,-1,6 == 7
	//0,3,7,2,5,8,4,6,0,1 == 9
	//100,4,200,1,3,2 ==4
	//0 == 1
	fmt.Println(longestConsecutive(nums))
}

func longestConsecutive(nums []int) int {
	myMap := map[int]bool{}

	for _, num := range nums {
		myMap[num] = true
	}

	count := 0

	for _, num := range nums {
		if myMap[num-1] {
			continue
		}

		tempcount := 1
		temp := num + 1

		for myMap[temp] {
			tempcount++
			temp++
		}

		if tempcount > count {
			count = tempcount
		}
	}

	return count
}