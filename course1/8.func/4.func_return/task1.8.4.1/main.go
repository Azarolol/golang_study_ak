package main

import "fmt"

func DivideAndRemainder(a, b int) (int, int) {
	defer func() {
		if r := recover(); r != nil {
			// обработка ошибки деления на ноль
			fmt.Println("check zero argument")
		}
	}()
	return a / b, a % b
}

func main() {
	div, rem := DivideAndRemainder(19, 3)
	fmt.Printf("Частное: %d, Остаток: %d\n", div, rem)
	div, rem = DivideAndRemainder(19, 0)
	fmt.Printf("Частное: %d, Остаток: %d\n", div, rem)
}
