package main

import "fmt"

func Shift(xs []int) (int, []int) {
	n := len(xs) - 1
	if n == -1 {
		return 0, []int{}
	}
	return xs[0], append([]int{xs[n]}, xs[0:n]...)
}

func main() {
	xs := []int{1, 2, 3, 4, 5}
	firstElement, shiftedSlice := Shift(xs)
	fmt.Println(firstElement) // Вывод: 1
	fmt.Println(shiftedSlice) // Вывод: [5 1 2 3 4]
}
