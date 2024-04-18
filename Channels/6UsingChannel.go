package main

import "fmt"

func main() {
	c := make(chan int)

	//	Send
	go foo(c)

	//Receive
	bars(c)

}

// Send
func foo(c chan<- int) {
	c <- 25
}

// Receive
func bars(c <-chan int) {

	fmt.Println(<-c)
}
