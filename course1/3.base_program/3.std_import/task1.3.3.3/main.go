package main

import (
	"fmt"
	"math"
)

func Sin(x float64) float64 {
	return math.Sin(x)
}

func Cos(x float64) float64 {
	return math.Cos(x)
}

func main() {
	// Пример вызова функции
	resultSin := Sin(0.0)
	resultCos := Cos(0.0)
	fmt.Println(resultSin) // 0
	fmt.Println(resultCos) // 1
}
