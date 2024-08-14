package main

import (
	"fmt"
	"math"
)

func CalculatePercentageChange(inintialValue, finalValue float64) float64 {
	return math.Round(((finalValue-inintialValue)/inintialValue)*10000) / 100
}

func main() {
	fmt.Print(CalculatePercentageChange(150, 200))
}
