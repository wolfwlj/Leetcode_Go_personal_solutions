package main

import "fmt"

func main() {

	//Problem : 
	//Given an integer array nums, return an array answer such 
	//that answer[i] is equal to the product of all the elements of nums except nums[i].

	//The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.
	
	//You must write an algorithm that runs in O(n) time and without using the division operation.


    // input : 
	var nums  = []int{2,3,5,0}
	//output should be [24,12,8,6]
	fmt.Println("function result   :   ", productExceptSelf(nums))
}

func productExceptSelf(nums []int) []int {

	answer := make([]int, len(nums))
	prefixarr := make([]int, len(nums))
	postfixarr := make([]int, len(nums))

	for i := range nums {
		if i -1 < 0 {
			prefixarr[i] = 1
		} else {
			prefixarr[i] = nums[i - 1] * prefixarr[i-1]

		}
	}
	for i := len(prefixarr) - 1; i >= 0; i-- {
		if i +1 >=len(nums) {
			postfixarr[i] = 1
		} else {
			postfixarr[i] = nums[i + 1] * postfixarr[i+1]
		}
	}

	for i := 0; i < len(nums); i++ {
		 
			answer[i] = prefixarr[i] * postfixarr[i]
	}

	return answer
}
