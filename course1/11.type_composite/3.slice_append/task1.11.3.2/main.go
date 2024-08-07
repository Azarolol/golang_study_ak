package main

import "fmt"

func appendInt(xs *[]int, x ...int) {
	*xs = append(*xs, x...)
}

func main() {
	xs := []int{0, 1, 2}
	appendInt(&xs, 3, 4, 5)
	fmt.Println(xs)
}
