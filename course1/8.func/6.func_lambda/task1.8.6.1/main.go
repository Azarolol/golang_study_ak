package main

import "fmt"

func Sum(a ...int) int {
	result := 0
	for _, x := range a {
		result += x
	}
	return result
}

func Mul(a ...int) int {
	if len(a) == 0 {
		return 0
	}
	result := a[0]
	for i := 1; i < len(a); i++ {
		result *= a[i]
	}
	return result
}

func Sub(a ...int) int {
	if len(a) == 0 {
		return 0
	}
	result := a[0]
	for i := 1; i < len(a); i++ {
		result -= a[i]
	}
	return result
}

func MathOperate(op func(a ...int) int, a ...int) int {
	return op(a...)
}

func main() {
	fmt.Println(MathOperate(Sum, 1, 1, 3))
	fmt.Println(MathOperate(Mul, 1, 7, 3))
	fmt.Println(MathOperate(Sub, 13, 2, 3))
}
