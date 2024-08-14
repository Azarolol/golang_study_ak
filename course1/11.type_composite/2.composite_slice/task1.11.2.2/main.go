package main

import "fmt"

func MaxDifference(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	min, max := numbers[0], numbers[0]
	for i := 1; i < len(numbers); i++ {
		if numbers[i] > max {
			max = numbers[i]
		}
		if numbers[i] < min {
			min = numbers[i]
		}
	}
	return max - min
}

func main() {
	fmt.Println(MaxDifference([]int{0, 1, 2, 3, 4, 5, 6}))
}
