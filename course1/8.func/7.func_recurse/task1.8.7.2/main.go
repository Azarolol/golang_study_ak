package main

import "fmt"

func Fibonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func main() {
	fmt.Println(Fibonacci(0))
	fmt.Println(Fibonacci(1))
	fmt.Println(Fibonacci(2))
	fmt.Println(Fibonacci(3))
	fmt.Println(Fibonacci(4))
	fmt.Println(Fibonacci(5))
	fmt.Println(Fibonacci(6))
}
