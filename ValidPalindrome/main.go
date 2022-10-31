package main

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)


func main() {
	// nums := []int{1,2,3,1}
	s := "A man, a plan, a canal: Panama"
	fmt.Println(isPalindrome(s))

}


func isPalindrome(s string) bool {

	string1 := strip(s)
	fmt.Println(string1)

	string2 := SpaceStringsBuilder(string1) 

   	string3 := strings.ToLower(string2)
	   fmt.Println(string3)

	string4 := Reverse(string3)
    fmt.Println(string4)

	return string3 == string4

	
	
}

//al deze helper functies heb ik niet zelf gemaakt, go heeft gewoon vaak geen standard library voor dit soort dingen.
func Reverse(s string) string {
    size := len(s)
    buf := make([]byte, size)
    for start := 0; start < size; {
        r, n := utf8.DecodeRuneInString(s[start:])
        start += n
        utf8.EncodeRune(buf[size-start:], r)
    }
    return string(buf)
}

func strip(s string) string {
    var result strings.Builder
    for i := 0; i < len(s); i++ {
        b := s[i]
        if ('a' <= b && b <= 'z') ||
            ('A' <= b && b <= 'Z') ||
            ('0' <= b && b <= '9') ||
            b == ' ' {
            result.WriteByte(b)
        }
    }
    return result.String()
}


func SpaceStringsBuilder(str string) string {
    var b strings.Builder
    b.Grow(len(str))
    for _, ch := range str {
        if !unicode.IsSpace(ch) {
            b.WriteRune(ch)
        }
    }
    return b.String()
}