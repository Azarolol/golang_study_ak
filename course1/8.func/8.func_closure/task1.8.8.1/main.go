package main

import "fmt"

func createCounter() func() int {
	counter := 0
	return func() int {
		counter++
		return counter
	}
}

func main() {
	counter := createCounter()
	fmt.Println(counter()) // 1
	fmt.Println(counter()) // 2
	fmt.Println(counter()) // 3
}
