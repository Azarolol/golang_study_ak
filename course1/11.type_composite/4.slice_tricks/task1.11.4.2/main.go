package main

import "fmt"

func InsertAfterIDX(xs []int, idx int, x ...int) []int {
	if idx < 0 || idx >= len(xs) {
		return []int{}
	}
	result := append(xs[0:idx+1], x...)
	return append(result, xs[idx+1:]...)
}

func main() {
	// Пример использования функции
	xs := []int{1, 2, 3, 4, 5}
	result := InsertAfterIDX(xs, 2, 6, 7, 8)
	fmt.Println(result) // Вывод: [1 2 3 6 7 8 4 5]
}
