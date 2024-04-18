package main

import (
	"fmt"
	"sync"
)

func workers(id int, tasks <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range tasks {
		fmt.Printf("Worker %d processing task %d\n", id, task)

	}

}

func main() {
	fmt.Println("This is 13 FansOut File")

	tasks := make(chan int)

	numWorkers := 3

	var wg sync.WaitGroup
	wg.Add(numWorkers)

	for i := 1; i < numWorkers; i++ {
		go worker(i, tasks, &wg)
	}

}
