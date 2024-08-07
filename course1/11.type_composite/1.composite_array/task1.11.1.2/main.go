package main

import (
	"fmt"
	"sort"
)

func sortDescInt(xs [8]int) [8]int {
	buf := []int{xs[0], xs[1], xs[2], xs[3], xs[4], xs[5], xs[6], xs[7]}
	sort.Slice(buf, func(i int, j int) bool {
		return buf[i] > buf[j]
	})
	return [8]int{buf[0], buf[1], buf[2], buf[3], buf[4], buf[5], buf[6], buf[7]}
}

func sortAscInt(xs [8]int) [8]int {
	buf := []int{xs[0], xs[1], xs[2], xs[3], xs[4], xs[5], xs[6], xs[7]}
	sort.Ints(buf)
	return [8]int{buf[0], buf[1], buf[2], buf[3], buf[4], buf[5], buf[6], buf[7]}
}

func sortDescFloat(xs [8]float64) [8]float64 {
	buf := []float64{xs[0], xs[1], xs[2], xs[3], xs[4], xs[5], xs[6], xs[7]}
	sort.Slice(buf, func(i int, j int) bool {
		return buf[i] > buf[j]
	})
	return [8]float64{buf[0], buf[1], buf[2], buf[3], buf[4], buf[5], buf[6], buf[7]}
}

func sortAscFloat(xs [8]float64) [8]float64 {
	buf := []float64{xs[0], xs[1], xs[2], xs[3], xs[4], xs[5], xs[6], xs[7]}
	sort.Float64s(buf)
	return [8]float64{buf[0], buf[1], buf[2], buf[3], buf[4], buf[5], buf[6], buf[7]}
}

func main() {
	intArr := [8]int{5, 2, 8, 1, 9, 3, 7, 4}
	floatArr := [8]float64{5.5, 2.2, 8.8, 1.1, 9.9, 3.3, 7.7, 4.4}

	sortedIntDesc := sortDescInt(intArr)
	sortedIntAsc := sortAscInt(intArr)
	sortedFloatDesc := sortDescFloat(floatArr)
	sortedFloatAsc := sortAscFloat(floatArr)

	fmt.Println("Sorted Int Array (Descending):", sortedIntDesc)
	fmt.Println("Sorted Int Array (Ascending):", sortedIntAsc)
	fmt.Println("Sorted Float Array (Descending):", sortedFloatDesc)
	fmt.Println("Sorted Float Arrat (Ascending):", sortedFloatAsc)
}
