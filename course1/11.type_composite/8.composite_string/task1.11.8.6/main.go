package main

import (
	"fmt"
	"strings"
)

func CountVowels(str string) int {
	vowels := "аеёиоуыэюяaeiouy"
	count := 0
	for _, r := range str {
		if strings.ContainsRune(vowels, r) {
			count++
		}
	}
	return count
}

func main() {
	// Пример использования функции CountVowels
	count := CountVowels("Привет, мир!")
	fmt.Println(count) // Вывод: 3
	count = CountVowels("Hello, world!")
	fmt.Println(count) // Вывод: 3
}
