package main

import "fmt"

func RemoveIDX(xs []int, idx int) []int {
	if idx < 0 || idx >= len(xs) {
		return xs
	}
	return append(xs[0:idx], xs[idx+1:]...)
}

func main() {
	xs := []int{0, 1, 2, 3, 4, 5}
	newXs := RemoveIDX(xs, 4)
	fmt.Println(newXs)
}
