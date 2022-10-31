package main

import (
	"fmt"
)



func main() {
	// nums := []int{1,2,3,1}
	nums := []int{8,9,2,5}
	fmt.Println(containsDuplicate(nums))
}

func containsDuplicate(nums []int) bool {
	var myMap = make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, check := myMap[nums[i]]; check{
			return true
		}
		myMap[nums[i]] = i


	}
	fmt.Println(myMap)
	return false
}
