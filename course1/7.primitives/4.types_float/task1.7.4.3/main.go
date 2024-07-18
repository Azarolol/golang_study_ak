package main

import (
	"fmt"
	"math"
)

func CompareRoundedValues(a, b float64, decimalPlaces int) (isEqual bool, difference float64) {
	rounder := math.Pow10(decimalPlaces)
	newA := math.Round(a*rounder) / rounder
	newB := math.Round(b*rounder) / rounder
	isEqual = newA == newB
	difference = math.Abs(newA - newB)
	return
}

func main() {
	fmt.Print(CompareRoundedValues(1.54443, 1.536213, 2))
}
