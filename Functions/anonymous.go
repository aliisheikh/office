package main

import "fmt"

func main() {
	fuo()

	func() {
		fmt.Println("Anonymous Function runs")
	}()

	func(a string) {
		fmt.Println("Hi! My name is ", a)
	}("Ali")

}
func fuo() {
	fmt.Println("Hi! I'm from fuo")

}
