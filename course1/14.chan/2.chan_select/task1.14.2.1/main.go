package main

import (
	"fmt"
	"time"
)

func trySend(ch chan int, v int) bool {
	select {
	case ch <- v:
		return true
	case <-time.After(2 * time.Second):
		return false
	}
}

func main() {
	ch := make(chan int, 1)
	defer close(ch)
	fmt.Println(trySend(ch, 1))
}
