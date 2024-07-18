package main

import (
	"fmt"
	"math"
	"math/big"
)

func Add(a, b int) *int {
	ans := a + b
	return &ans
}

func Max(numbers []int) *int {
	max := math.MinInt
	for _, num := range numbers {
		if num > max {
			max = num
		}
	}
	return &max
}

func IsPrime(number int) *bool {
	ans := big.NewInt(int64(number)).ProbablyPrime(0)
	return &ans
}

func ConcatenateStrings(strs []string) *string {
	ans := ""
	for _, str := range strs {
		ans += str
	}
	return &ans
}

func main() {
	fmt.Println(*Add(2, 3))
	fmt.Println(*Max([]int{1, 2, 3}))
	fmt.Println(*IsPrime(1212121))
	fmt.Println(*ConcatenateStrings([]string{"Hello", ", ", "world", "!"}))
}
