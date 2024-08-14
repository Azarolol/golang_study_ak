package main

import (
	"fmt"
	"time"

	"github.com/mattevans/dinero"
)

const (
	DINERO_API = "40060c7df7364e15913ce40a7a3b4e6b"
)

func currencyPairRate(a, b string, x float64) float64 {
	client := dinero.NewClient(DINERO_API, a, 20*time.Minute)
	rate, _ := client.Rates.Get(b)
	return *rate * x
}

func main() {
	fmt.Println(currencyPairRate("USD", "EUR", 100))
}
