package main

import "fmt"

func main() {

	channel := make(chan string)

	go func() {
		channel <- "Muhammad Ali"
	}()
	fmt.Printf("The Type of channel is %T\n", channel)
	fmt.Println(<-channel)

}
