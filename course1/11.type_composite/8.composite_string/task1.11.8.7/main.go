package main

import (
	"fmt"
	"strings"
)

func ReplaceSymbols(str string, old, new rune) string {
	return strings.ReplaceAll(str, string(old), string(new))
}

func main() {
	// Пример использования функции ReplaceSymbols
	result := ReplaceSymbols("Hello, world", 'o', '0')
	// result должно быть равно "Hell0, w0rld!"
	fmt.Println(result)
}
