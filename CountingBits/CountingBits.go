package main

import (
	"fmt"
	"strings"
)

func main() {

	var n int = 5

	fmt.Println(countBits(n))
}

func countBits(n int) []int {
	var ans = []int{}

	for i := 0; i <= n; i++ {
		substr := "1"
		convToBinary := int64(i)
		bin := fmt.Sprintf("%b", convToBinary)

		count := strings.Count(bin, substr)

		fmt.Println(i, " | ", bin, "  |  ", count)
		// result :
		//(index) (binary) (aantal 1 in binary)
		// 0  	|  0   	|   0
		// 1  	|  1   	|   1
		// 2  	|  10   |   1
		// 3  	|  11   |   2
		// 4  	|  100  |   1
		// 5  	|  101  |   2
		ans = append(ans, count)
		// ans test case output = [0 1 1 2 1 2]
	}

	return ans
}
