package main

import "fmt"

func FindMaxAndMin(n ...int) (int, int) {
	if len(n) == 0 {
		return 0, 0
	}
	max := n[0]
	min := n[0]
	for i := 1; i < len(n); i++ {
		if n[i] > max {
			max = n[i]
		}
		if n[i] < min {
			min = n[i]
		}
	}
	return max, min
}

func main() {
	fmt.Print(FindMaxAndMin(1, 2, 3, 4, 5))
}
