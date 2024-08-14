package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	value int
	mutex sync.Mutex
}

func (c *Counter) Increment() int {
	c.mutex.Lock()
	c.value++
	c.mutex.Unlock()
	return c.value
}

func concurrentSafeCounter() int {
	wg := sync.WaitGroup{}
	counter := Counter{}
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}
	wg.Wait()
	return counter.value
}

func main() {
	fmt.Println(concurrentSafeCounter())
}
