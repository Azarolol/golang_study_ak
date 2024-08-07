package main

import (
	"fmt"
	"sync"
)

type Worker struct {
	ID     int
	Job    chan int
	Result chan int
}

func work(w Worker, wg *sync.WaitGroup) {
	for job := range w.Job {
		// Perform the task associated with the job.
		result := job * 2

		// Send the result to the result channel.
		w.Result <- result

		fmt.Printf("Job №%d done by worker №%d\n", job, w.ID)
	}
	fmt.Printf("Worker №%d is done\n", w.ID)
	wg.Done()
}

const numWorkers = 3

func main() {
	var wg sync.WaitGroup
	jobs := make(chan int)
	results := make(chan int)

	// Create the worker pool.
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		worker := Worker{
			ID:     i + 1,
			Job:    jobs,
			Result: results,
		}
		go work(worker, &wg)
	}

	numbers := []int{1, 2, 3, 4, 5}
	// Submit jobs to the worker pool.
	go func() {
		for _, job := range numbers {
			jobs <- job
		}
		// Close the job channel to indicate that no more jobs will be sent.
		close(jobs)
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect the results from the result channel.
	for result := range results {
		fmt.Println("Result:", result)
	}
}
