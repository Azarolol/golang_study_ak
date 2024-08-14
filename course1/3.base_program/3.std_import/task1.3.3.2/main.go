package main

import (
	"fmt"
	"math"
)

func Floor(x float64) float64 {
	return math.Floor(x)
}

func main() {
	// Пример вызова функции
	result := Floor(2.8)
	fmt.Println(result) // 2
}
