package main

import "fmt"

func PrintNumbers(nums ...int) {
	for _, num := range nums {
		fmt.Println(num)
	}
}

func main() {
	PrintNumbers(1, 2, 3, 4, 5)
}
