package main

import "fmt"

func Pop(xs []int) (int, []int) {
	if len(xs) == 0 {
		return 0, xs
	}
	if len(xs) == 1 {
		return xs[0], []int{}
	}
	return xs[0], xs[1:]
}

func main() {
	xs := []int{1, 2, 3, 4, 5}
	firstElement, newXs := Pop(xs)
	fmt.Printf("Значение: %d, Новый срез: %d", firstElement, newXs)
}
