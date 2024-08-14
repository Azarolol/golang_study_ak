package main

import (
	"fmt"
)

// Функция getBytes возвращает срез байтов строки s
func getBytes(s string) []byte {
	return []byte(s)
}

// Функция getRunes возвращает срез символов строки s
func getRunes(s string) []rune {
	return []rune(s)
}

func main() {
	// Пример использования функции getBytes
	bytes := getBytes("Привет, мир!")
	fmt.Println(bytes)

	// Пример использования функции getRunes
	runes := getRunes("Привет, мир!")
	fmt.Println(runes)
}
