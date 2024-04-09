package main

import "fmt"

func main() {

	// Defer just defer the function

	defer fou()
	bars()

}

func fou() {
	fmt.Println("I'm foo")

}
func bars() {
	fmt.Println("I'm bar")

}
