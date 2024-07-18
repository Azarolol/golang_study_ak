package main

import "fmt"

func CalculateStockValue(price float64, quantity int) (float64, float64) {
	return price * float64(quantity), price
}

func main() {
	fmt.Print(CalculateStockValue(1.0, 10))
}
