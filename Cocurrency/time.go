package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// Concurrency: Goroutines executing concurrently

	// Define a wait group to synchronize the completion of goroutines
	var wg sync.WaitGroup

	// Function to simulate a task that takes some time to complete
	task := func(id int) {
		defer wg.Done()         // Mark the task as done when finished
		time.Sleep(time.Second) // Simulate work by sleeping for 1 second
		fmt.Printf("Goroutine %d has finished.\n", id)
	}

	// Create and start multiple goroutines concurrently
	for i := 1; i <= 5; i++ {
		wg.Add(1)  // Increment the wait group counter
		go task(i) // Start a new goroutine to execute the task
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Parallelism: Goroutines executing in parallel

	// Set the number of OS threads to use
	// This enables parallelism by allowing multiple goroutines to run concurrently on different OS threads
	// By default, GOMAXPROCS is set to the number of CPU cores
	// For demonstration purposes, we'll set it to 2 to limit parallelism
	// Uncomment the line below to observe parallelism
	// runtime.GOMAXPROCS(2)

	// Define another wait group for parallel tasks
	var pwg sync.WaitGroup

	// Parallel task function
	ptask := func(id int) {
		defer pwg.Done()
		time.Sleep(time.Second)
		fmt.Printf("Parallel task %d has finished.\n", id)
	}

	// Start two tasks in parallel
	pwg.Add(2)
	go ptask(1)
	go ptask(2)

	// Wait for parallel tasks to finish
	pwg.Wait()

	fmt.Println("All tasks have finished.")
}
