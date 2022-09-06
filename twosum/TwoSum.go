package main

import "fmt"

func main() {

	var nums = []int{2, 7, 11, 15}
	var target int = 9

	fmt.Println(Twosum(nums, target))
}

func Twosum(nums []int, target int) []int {
	//Langzame oplossing
	//var arr = []int{}

	// for i := 0; i < len(nums); i++ {
	// 	for j := i + 1; j < len(nums); j++ {
	// 		if nums[i]+nums[j] == target {
	// 			arr = append(arr, i, j)
	// 			return arr
	// 		}
	// 	}
	// }

	//twosum
	// var arr = []int{}

	//hash map oplossing
	var arr = []int{}
	//hash map de map[int]int betekent, de linker int is de key en de rechter int is de value en je verwacht dat ze beide ints zijn
	var myMap = make(map[int]int)

	for i := 0; i < len(nums); i++ {

		//check of myMap[target-nums[i]] in de hash map zit
		if _, ok := myMap[target-nums[i]]; ok {
			//zoja dat betekent dat het target bereikt is
			arr = append(arr, myMap[target-nums[i]], i)
			// zet de index van de hash map in de array
			// return de array
			return arr
		}
		//als het target nog niet bereikt is voeg de waarde nums[i] en de index toe aan de hash map
		myMap[nums[i]] = i

	}
	// deze return is alleen als er geen target bereikt is, en er is iets mis gegaan
	return nums
}
