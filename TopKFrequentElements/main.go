package main

import (
	"fmt"
	"sort"
)

func main() {

	var nums = []int{3,3,1,2,2,3}
	var k int = 2

	fmt.Print(topKFrequent(nums, k))
}

func topKFrequent(nums []int, k int) []int {
	answer := []int{}

	var myMap = make(map[int]int)

	for i := 0; i < len(nums); i++ {
		myMap[nums[i]] = myMap[nums[i]] + 1
	}

	// fmt.Print(myMap)
    keys := make([]int, 0, len(myMap))
	fmt.Print(keys)
 
    for key := range myMap {
        keys = append(keys, key)
    }
	fmt.Println(keys)
 
    sort.SliceStable(keys, func(i, j int) bool{
        return myMap[keys[i]] > myMap[keys[j]]
    })
	fmt.Println(keys)
 
	for j := 0; j < k; j++ {

		answer = append(answer, keys[j])

	}
	return answer
}
