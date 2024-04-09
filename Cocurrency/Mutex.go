package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("GOr outine", runtime.NumGoroutine())

	counter := 0

	const sd = 100
	var wg sync.WaitGroup
	wg.Add(sd)

	var mu sync.Mutex

	for i := 0; i < sd; i++ {
		go func() {
			mu.Lock()
			v := counter
			v++
			counter = v
			go runtime.Gosched()
		}()
		fmt.Println("GO Routine", runtime.NumGoroutine())

		wg.Done()
		mu.Unlock()

	}
	wg.Wait()

	fmt.Println("GO routine", runtime.NumGoroutine())
	fmt.Println("Counter", counter)

}

// We Use Mutex to Lock down certain chunks of code So that Multiple
//Go routine can't Access that code at the same time.
