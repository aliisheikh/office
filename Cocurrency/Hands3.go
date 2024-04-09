package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	fmt.Println("CPUs", runtime.NumCPU())
	fmt.Println("GO routines", runtime.NumGoroutine())

	count := 0
	//var mu sync.Mutex

	var wg sync.WaitGroup
	a := 100
	wg.Add(a)

	for i := 1; i < a; i++ {
		go func() {
			//mu.Lock()
			v := count
			runtime.Gosched()
			v++
			count = v

			fmt.Println(count)
			//mu.Unlock()
			wg.Done()

		}()

	}
	wg.Wait()

	fmt.Println("Count is ", count)

}
