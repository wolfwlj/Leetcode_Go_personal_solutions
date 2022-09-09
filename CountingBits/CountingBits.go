package main

import (
	"fmt"
)

func main() {

	var n int = 8

	fmt.Println(countBits(n))
	countBits(n)
}

func countBits(n int) []int {
	// var ans = []int{}

	// for i := 0; i <= n; i++ {
	// 	substr := "1"
	// 	convToBinary := int64(i)
	// 	bin := fmt.Sprintf("%b", convToBinary)

	// 	count := strings.Count(bin, substr)

	// 	fmt.Println(i, " | ", bin, "  |  ", count)

	// 	// result :
	// 	//(index) (binary) (aantal 1 in binary)
	// 	// 0  	|  0   	|   0
	// 	// 1  	|  1   	|   1
	// 	// 2  	|  10   |   1
	// 	// 3  	|  11   |   2
	// 	// 4  	|  100  |   1
	// 	// 5  	|  101  |   2
	// 	ans = append(ans, count)
	// 	// ans test case output = [0 1 1 2 1 2]
	// }

	//0000
	//0001
	//0010
	//0011

	//0100
	//0101
	//0110
	//0111

	//1000
	// var length int = n + 1
	/// veel betere oplossing.
	//reken uit wanneer je bij een nieuw macht van 2 komt, daarna kijk je naar het current nummer - de laatste macht.
	var ans = []int{0}
	var offset int = 1

	for i := 1; i <= n; i++ {

		if offset*2 == i {
			offset = i
		}
		ans = append(ans, 1+ans[i-offset])
	}
	//return final array
	return ans
}
