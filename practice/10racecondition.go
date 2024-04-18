package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	fmt.Println("CPUs", runtime.NumCPU())
	fmt.Println("Go routines", runtime.NumGoroutine())
	count := 0
	const gr = 100

	var wg sync.WaitGroup
	wg.Add(gr)

	//var mu sync.Mutex

	for i := 0; i < gr; i++ {
		go func() {

			//mu.Lock()
			v := count
			runtime.Gosched()
			v++
			count = v
			//mu.Unlock()
			wg.Done()
		}()
		fmt.Println("Go routines", runtime.NumGoroutine())

	}
	wg.Wait()
	fmt.Println("Goroutine", runtime.NumGoroutine())

	fmt.Println("count", count)

}
