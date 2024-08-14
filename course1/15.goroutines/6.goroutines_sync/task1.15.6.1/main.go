package main

import (
	"fmt"
	"sync"
)

func waitGroupExample(goroutines ...func() string) string {
	mutex := sync.Mutex{}
	wg := sync.WaitGroup{}
	result := ""
	wg.Add(len(goroutines))
	for _, f := range goroutines {
		go func(f func() string) {
			defer wg.Done()
			mutex.Lock()
			result += f() + " done\n"
			mutex.Unlock()
		}(f)
	}
	wg.Wait()
	return result
}

func main() {
	count := 1000
	goroutines := make([]func() string, count)
	for i := 0; i < count; i++ {
		j := i
		goroutines[i] = func() string {
			return fmt.Sprintf("goroutine %d", j)
		}
	}

	fmt.Println(waitGroupExample(goroutines...))
}
