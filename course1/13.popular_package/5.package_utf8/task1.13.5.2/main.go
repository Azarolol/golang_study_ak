package main

import (
	"fmt"
	"unicode/utf8"
)

func countRussianLetters(s string) map[rune]int {
	counts := make(map[rune]int)
	for _, char := range s {
		if isRussianLetter(char) && utf8.ValidRune(char) {
			if char < 1072 {
				char += 32
			}
			if _, ok := counts[char]; !ok {
				counts[char] = 1
			} else {
				counts[char]++
			}
		}
	}
	return counts
}

func isRussianLetter(r rune) bool {
	return r > 1039 && r < 1104
}

func main() {
	result := countRussianLetters("Привет, мир!")
	for key, value := range result {
		fmt.Printf("%c: %d ", key, value) // в: 1 е: 1 т: 1 м: 1 п: 1 р: 2 и: 2
	}
}
