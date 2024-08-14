package main

import (
	"fmt"
	"unicode/utf8"
)

func countBytes(s string) int {
	return len([]byte(s))
}

func countSymbols(s string) int {
	return utf8.RuneCountInString(s)
}

func main() {
	// Пример использования функции countBytes
	bytes := countBytes("Привет, мир!")
	fmt.Println(bytes) // Вывод: 21

	// Пример использования функции countSymbols
	symbols := countSymbols("Привет, мир!")
	fmt.Println(symbols) // Вывод: 12
}
