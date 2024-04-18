package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {

	fmt.Println("CPUs ->", runtime.NumCPU())
	fmt.Println("Go-routines", runtime.NumGoroutine())

	var count int64
	const gr = 100

	var wg sync.WaitGroup
	wg.Add(gr)

	for i := 0; i < gr; i++ {

		go func() {
			atomic.AddInt64(&count, 1)
			runtime.Gosched()
			fmt.Println(atomic.LoadInt64(&count))
			wg.Done()
		}()
		fmt.Println("Go-routine:", runtime.NumGoroutine())

	}
	wg.Wait()
	fmt.Println("Go-routine:", runtime.NumGoroutine())
	fmt.Println("Counter:", count)

}
