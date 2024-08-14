package main

import (
	"fmt"
	"sync"
)

func mergeChan2(from ...chan int) chan int {
	var wg sync.WaitGroup

	merged := make(chan int, 100)

	wg.Add(len(from))

	output := func(ch chan int) {
		for x := range ch {
			merged <- x
		}
		wg.Done()
	}

	for _, ch := range from {
		go output(ch)
	}

	go func() {
		wg.Wait()
		close(merged)
	}()

	return merged
}

func mergeChan(mergeTo chan int, from ...chan int) {
	var wg sync.WaitGroup

	wg.Add(len(from))

	output := func(ch chan int) {
		for x := range ch {
			mergeTo <- x
		}
		wg.Done()
	}

	for _, ch := range from {
		go output(ch)
	}

	go func() {
		wg.Wait()
		close(mergeTo)
	}()
}

func generateChan(n int) chan int {
	ch := make(chan int)

	go func() {
		for i := 0; i < n; i++ {
			ch <- i
		}
		close(ch)
	}()
	return ch
}

func main() {
	chanInputNums1 := generateChan(101)
	chanInputNums2 := generateChan(101)
	chanInputNums3 := generateChan(101)

	ch := make(chan int, 100)
	mergeChan(ch, chanInputNums1, chanInputNums2)
	merged := mergeChan2(ch, chanInputNums3)
	sum := 0

	for num := range merged {
		sum += num
	}

	fmt.Println("Sum is", sum, "should be 15150")
}
