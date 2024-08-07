package main

import "fmt"

func Cut(xs []int, start, end int) []int {
	if start < 0 || start > end || end > len(xs) {
		return []int{}
	}
	return xs[start : end+1]
}

func main() {
	xs := []int{1, 2, 3, 4, 5}
	result := Cut(xs, 1, 3)
	fmt.Println(result) // Вывод: [2, 3, 4]
}
