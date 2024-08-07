package main

import "fmt"

func InsertToStart(xs []int, x ...int) []int {
	return append(x, xs...)
}

func main() {
	xs := []int{1, 2, 3}
	result := InsertToStart(xs, 4, 5, 6)
	fmt.Println(result) // {4, 5, 6, 1, 2, 3}
}
