package main

import (
	"fmt"
	"sync"
)

// Worker function that processes tasks
func worker(id int, tasks <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range tasks {
		fmt.Printf("Worker %d processing task %d\n", id, task)
		// Perform some computation or task processing
	}
}

func main() {
	// Create an input channel for tasks
	tasks := make(chan int)

	// Define the number of worker goroutines
	numWorkers := 3

	// Create a wait group to wait for all worker goroutines to finish
	var wg sync.WaitGroup
	wg.Add(numWorkers)

	// Spawn worker goroutines
	for i := 1; i <= numWorkers; i++ {
		go worker(i, tasks, &wg)
	}

	// Generate some tasks and distribute them to the workers
	for i := 1; i <= 10; i++ {
		tasks <- i
	}

	// Close the tasks channel to signal that no more tasks will be sent
	close(tasks)

	// Wait for all worker goroutines to finish
	wg.Wait()
}
