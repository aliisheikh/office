package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	//fmt.Println("CPUs", runtime.NumCPU())
	//fmt.Println("GO routines", runtime.NumGoroutine())
	var wg sync.WaitGroup
	var count int64

	a := 100
	wg.Add(a)

	for i := 0; i < a; i++ {
		go func() {

			atomic.AddInt64(&count, 1)
			fmt.Println(atomic.LoadInt64(&count))

			wg.Done()

		}()

	}
	wg.Wait()

	fmt.Println("Count is ", count)

}
