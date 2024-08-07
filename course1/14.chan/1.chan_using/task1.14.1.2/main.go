package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateData(n int) chan int {
	rand.Seed(time.Now().UnixNano())
	ch := make(chan int, 10)
	go func() {
		for i := 0; i < n; i++ {
			ch <- rand.Int()
		}
		close(ch)
	}()
	return ch
}

func main() {
	data := generateData(10)
	go func() {
		time.Sleep(1 * time.Second)
		close(data)
	}()

	for num := range data {
		fmt.Println(num)
	}
}
