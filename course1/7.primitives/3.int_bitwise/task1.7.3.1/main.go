package main

import "fmt"

func bitwiseAnd(a, b int) int {
	return a & b
}

func bitwiseOr(a, b int) int {
	return a | b
}

func bitwiseXor(a, b int) int {
	return a ^ b
}

func bitwiseLeftShift(a, b int) int {
	return a << b
}

func bitwiseRightShift(a, b int) int {
	return a >> b
}

func main() {
	a := 5
	b := 1
	fmt.Printf("a & b = %d\n", bitwiseAnd(a, b))
	fmt.Printf("a | b = %d\n", bitwiseOr(a, b))
	fmt.Printf("a ^ b = %d\n", bitwiseXor(a, b))
	fmt.Printf("a << b = %d\n", bitwiseLeftShift(a, b))
	fmt.Printf("a >> b = %d\n", bitwiseRightShift(a, b))
}
