package main

import (
	"fmt"
)

func sum(xs [8]int) int {
	result := 0
	for _, x := range xs {
		result += x
	}
	return result
}

func average(xs [8]int) float64 {
	return float64(sum(xs)) / 8
}

func averageFloat(ys [8]float64) float64 {
	result := 0.0
	for _, y := range ys {
		result += y
	}
	return result / 8
}

func reverse(xs [8]int) [8]int {
	return [8]int{xs[7], xs[6], xs[5], xs[4], xs[3], xs[2], xs[1], xs[0]}
}

func main() {
	xs := [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(sum(xs)) // Вывод: 36

	fmt.Println(average(xs)) // Вывод: 4.5

	ys := [8]float64{1.5, 2.5, 3.5, 4.5, 5.5, 6.5, 7.5, 8.5}

	fmt.Println(averageFloat(ys)) // Вывод: 5

	fmt.Println(reverse(xs)) // Вывод: [8 7 6 5 4 3 2 1]
}
