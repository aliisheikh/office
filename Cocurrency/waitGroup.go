package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {

	fmt.Println("OS \t", runtime.GOOS)
	fmt.Println("ARCH", runtime.GOARCH)
	fmt.Println("CPUs", runtime.NumCPU())
	fmt.Println("GO routines", runtime.NumGoroutine())

	wg.Add(1)
	go foo()
	bar()
	fmt.Println("CPUs", runtime.NumCPU())
	fmt.Println("GO routines", runtime.NumGoroutine())
	wg.Wait()

}

func foo() {

	for i := 1; i < 10; i++ {
		fmt.Println("Foo Function", i)

	}
	wg.Done()

}
func bar() {
	for i := 1; i < 10; i++ {
		fmt.Println("Bar Function", i)
	}

}
