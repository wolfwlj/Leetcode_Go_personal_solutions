package main

import (
	"fmt"
)

func main() {

	var s string = "s"
	var t string = ""

	fmt.Println(isSubsequence(s, t))
	// fmt.Println(len(s))
	// fmt.Println(len(s))

}

func isSubsequence(s string, t string) bool {
	i, j := 0, 0
	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			i += 1
		}
		j += 1

	}
	return i == len(s)
}
