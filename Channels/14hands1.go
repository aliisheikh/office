package main

import "fmt"

func main() {

	//With Buffer
	c := make(chan int, 1)

	c <- 20
	// With Anonymous Function
	//go func() {
	//	c <- 20
	//}()

	fmt.Println("The Value of c is", <-c)
}
