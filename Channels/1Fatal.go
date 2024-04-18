package main

import "fmt"

func main() {

	c := make(chan int)
	// This can't be run without Anonymous Function
	// Because of Channels Block.
	c <- 40

	fmt.Println(<-c)

}
