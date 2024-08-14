package main

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func DecimalSum(a, b string) (string, error) {
	firstDecimal, err := decimal.NewFromString(a)
	if err != nil {
		return "", err
	}
	secondDecimal, err := decimal.NewFromString(b)
	if err != nil {
		return "", err
	}
	sum := firstDecimal.Add(secondDecimal)
	return sum.String(), nil
}

func DecimalSubtract(a, b string) (string, error) {
	firstDecimal, err := decimal.NewFromString(a)
	if err != nil {
		return "", err
	}
	secondDecimal, err := decimal.NewFromString(b)
	if err != nil {
		return "", err
	}
	sub := firstDecimal.Sub(secondDecimal)
	return sub.String(), nil
}

func DecimalMultiply(a, b string) (string, error) {
	firstDecimal, err := decimal.NewFromString(a)
	if err != nil {
		return "", err
	}
	secondDecimal, err := decimal.NewFromString(b)
	if err != nil {
		return "", err
	}
	mul := firstDecimal.Mul(secondDecimal)
	return mul.String(), nil
}

func DecimalDivide(a, b string) (string, error) {
	firstDecimal, err := decimal.NewFromString(a)
	if err != nil {
		return "", err
	}
	secondDecimal, err := decimal.NewFromString(b)
	if err != nil {
		return "", err
	}
	if secondDecimal.Equal(decimal.Zero) {
		return "", fmt.Errorf("Division by 0")
	}
	div := firstDecimal.Div(secondDecimal)
	return div.String(), nil
}

func DecimalRound(a string, precision int32) (string, error) {
	dec, err := decimal.NewFromString(a)
	if err != nil {
		return "", err
	}
	round := dec.Round(precision)
	return round.String(), nil
}

func DecimalGreaterThan(a, b string) (bool, error) {
	firstDecimal, err := decimal.NewFromString(a)
	if err != nil {
		return false, err
	}
	secondDecimal, err := decimal.NewFromString(b)
	if err != nil {
		return false, err
	}
	return firstDecimal.GreaterThan(secondDecimal), nil
}

func DecimalLessThan(a, b string) (bool, error) {
	firstDecimal, err := decimal.NewFromString(a)
	if err != nil {
		return false, err
	}
	secondDecimal, err := decimal.NewFromString(b)
	if err != nil {
		return false, err
	}
	return firstDecimal.LessThan(secondDecimal), nil
}

func DecimalEqual(a, b string) (bool, error) {
	firstDecimal, err := decimal.NewFromString(a)
	if err != nil {
		return false, err
	}
	secondDecimal, err := decimal.NewFromString(b)
	if err != nil {
		return false, err
	}
	return firstDecimal.Equal(secondDecimal), nil
}

func main() {

}
