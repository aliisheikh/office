package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {

	fmt.Println("CPU", runtime.NumCPU())
	fmt.Println("Goroutine", runtime.NumGoroutine())
	//int 64 represents Atomic
	var counter int64

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			fmt.Println("Counter:", atomic.LoadInt64(&counter))

			runtime.Gosched()
			wg.Done()

		}()
		fmt.Println("Go Routine", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Go Routines", runtime.NumGoroutine())
	fmt.Println("Count:", counter)
}

//Package Atomic is used to write and read from int 64 and
//also Avoid Race Condition
