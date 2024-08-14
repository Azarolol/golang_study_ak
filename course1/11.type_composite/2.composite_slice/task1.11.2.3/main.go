package main

import "fmt"

func bitwiseXOR(n, res int) int {
	return n ^ res
}

func findSingleNumber(numbers []int) int {
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers); j++ {
			if i != j && bitwiseXOR(numbers[i], numbers[j]) == 0 {
				break
			}
			if j == len(numbers)-1 {
				return numbers[i]
			}
		}
	}
	return 0
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 4, 3, 2, 1}
	singleNumber := findSingleNumber(numbers)
	fmt.Println(singleNumber) // 5
}
