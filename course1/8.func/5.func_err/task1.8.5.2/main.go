package main

import "fmt"

func CheckDiscount(price, discount float64) (float64, error) {
	if discount <= 50 {
		return price * (100 - discount) / 100, nil
	}
	return 0, fmt.Errorf("скидка не может превышать 50%%")
}

func main() {
	fmt.Println(CheckDiscount(1500, 49))
	fmt.Println(CheckDiscount(1500, 51))
}
