package main

import (
	"fmt"
	"unicode/utf8"
)

func countUniqueUTF8Chars(s string) int {
	dict := make(map[rune]struct{})
	result := 0
	for _, r := range s {
		if utf8.ValidRune(r) {
			if _, ok := dict[r]; !ok {
				result++
				dict[r] = struct{}{}
			}
		}
	}
	return result
}

func main() {
	fmt.Println(countUniqueUTF8Chars("Hel1o, 	!")) // 9
}
