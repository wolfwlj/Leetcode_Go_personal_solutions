package main

import (
	"fmt"
)


func main() {
	// s := "anagram" 
	// t := "nagaram"
	// s := "rat" 
	// t := "car"
	s := "aacc"
	t := "ccac"

	fmt.Println(isAnagram(s, t))
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t){return false}
	var countS = make(map[byte]int)
	var countT = make(map[byte]int)

	for i := 0; i < len(s); i++ {

		countS[s[i]] = 1 + countS[s[i]]
		countT[t[i]] = 1 + countT[t[i]]

	}	
	fmt.Println(countS)
	fmt.Println(countT)


	for i := range countS {
		if countS[i] != countT[i] {
			return false
		}

	}
	
	return true
}