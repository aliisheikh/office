package main

import (
	"fmt"
	"sync"
)

func main() {

	//fmt.Println("CPUs", runtime.NumCPU())
	//fmt.Println("GO routines", runtime.NumGoroutine())
	var wg sync.WaitGroup
	count := 0
	var mu sync.Mutex

	a := 100
	wg.Add(a)

	for i := 0; i < a; i++ {
		go func() {
			mu.Lock()
			v := count
			v++
			count = v

			fmt.Println(count)
			mu.Unlock()
			wg.Done()

		}()

	}
	wg.Wait()

	fmt.Println("Count is ", count)

}
