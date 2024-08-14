package main

import "fmt"

func RemoveExtraMemory(xs []int) []int {
	newXs := make([]int, len(xs))
	copy(newXs, xs)
	return newXs
}

func main() {
	xs := make([]int, 3, 10)
	xs[0] = 1
	xs[1] = 2
	xs[2] = 3
	newXs := RemoveExtraMemory(xs)
	fmt.Println(xs, len(xs), cap(xs))
	fmt.Println(newXs, len(newXs), cap(newXs))
}
