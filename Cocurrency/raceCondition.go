package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("GO routine:", runtime.NumGoroutine())

	counter := 0

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {

			v := counter
			runtime.Gosched()
			v++
			counter = v

			wg.Done()
		}()
		fmt.Println("Goroutine", runtime.NumGoroutine())

	}
	wg.Wait()
	fmt.Println("Go routine", runtime.NumGoroutine())
	fmt.Println("Counter:", counter)

}
