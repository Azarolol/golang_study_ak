package main

import "fmt"

func ReverseString(str string) string {
	runes := []rune(str)
	for i := 0; i < len(runes)/2; i++ {
		runes[i], runes[len(runes)-i-1] = runes[len(runes)-i-1], runes[i]
	}
	return string(runes)
}

func main() {
	fmt.Println(ReverseString("Hello, world!")) // Вывод: "!dlrow ,olleH"
	fmt.Println(ReverseString("12345"))         // Вывод: "54321"
}
