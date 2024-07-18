package main

import (
	"fmt"
	"math"
	"strconv"
)

func CalculatePercentageChange(inintialValue, finalValue string) (float64, error) {
	initialFloat, err := strconv.ParseFloat(inintialValue, 64)
	if err != nil {
		return 0, err
	}
	finalFloat, err := strconv.ParseFloat(finalValue, 64)
	if err != nil {
		return 0, err
	}
	return math.Round(((finalFloat-initialFloat)/initialFloat)*10000) / 100, nil
}

func main() {
	fmt.Println(CalculatePercentageChange("150", "200"))
	fmt.Println(CalculatePercentageChange("150g", "200"))
}
