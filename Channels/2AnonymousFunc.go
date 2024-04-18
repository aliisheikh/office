package main

import "fmt"

func main() {
	c := make(chan int)

	//	Here We Call Anonymous Function

	go func() {
		c <- 50
	}()
	fmt.Println("The Value of <-c is:", <-c)
}
