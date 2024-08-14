package main

import "fmt"

func appendInt(xs []int, x ...int) []int {
	return append(xs, x...)
}

func main() {
	fmt.Println(appendInt([]int{0, 1, 2}, 3, 4, 5, 6, 7))
}
