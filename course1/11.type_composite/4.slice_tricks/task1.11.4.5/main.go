package main

import "fmt"

func FilterDividers(xs []int, divider int) []int {
	if divider == 0 {
		return []int{}
	}
	newXs := make([]int, 0)
	for _, x := range xs {
		if x%divider == 0 {
			newXs = append(newXs, x)
		}
	}
	return newXs
}

func main() {
	xs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	newXs := FilterDividers(xs, 2)
	fmt.Println(newXs)
}
