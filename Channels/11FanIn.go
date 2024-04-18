package main

import "fmt"

func main() {

	fmt.Println("This is FanIn File")
	// Create input channels
	ch1 := make(chan int)
	ch2 := make(chan int)

	// Start producing data on input channels
	go func() {
		defer close(ch1)
		for i := 0; i < 5; i++ {
			ch1 <- i * 2
		}
	}()

	go func() {
		defer close(ch2)
		for i := 0; i < 5; i++ {
			ch2 <- i * 3
		}
	}()

	// Combine data from input channels into a single output channel
	output := fanIn(ch1, ch2)

	// Consume data from the output channel
	for v := range output {
		fmt.Println(v)
	}
}

// Fan-in function takes multiple input channels and combines their data into a single output channel
func fanIn(input1, input2 <-chan int) <-chan int {
	output := make(chan int)

	go func() {
		defer close(output)
		for {
			select {
			case v := <-input1:
				output <- v
			case v := <-input2:
				output <- v
			}
		}
	}()

	return output
}
