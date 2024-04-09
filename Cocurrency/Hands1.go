package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	fmt.Println("Begin CPUs", runtime.NumCPU())
	fmt.Println("OS", runtime.GOOS)
	fmt.Println("ARCH", runtime.GOARCH)
	fmt.Println("Begins Go routine", runtime.NumGoroutine())

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		fmt.Println("Hello From Func one")
		wg.Done()
	}()

	go func() {
		fmt.Println("Hello From Func Two")
		wg.Done()
	}()
	fmt.Println("Mid CPU", runtime.NumCPU())
	fmt.Println("Mid Go Routine", runtime.NumGoroutine())

	wg.Wait()

	fmt.Println("Exit ")
	fmt.Println("End CPU", runtime.NumCPU())
	fmt.Println("End Go Routine", runtime.NumGoroutine())

}
