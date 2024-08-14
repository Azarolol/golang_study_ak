package main

import "fmt"

func adder(initial int) func(int) int {
	return func(i int) int {
		return initial + i
	}
}

func main() {
	// пример использования функции adder
	addTwo := adder(2)
	result := addTwo(3)
	fmt.Println(result) // выводит 5
}
