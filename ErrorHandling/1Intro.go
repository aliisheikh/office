package main

import "fmt"

func main() {

	fmt.Println("This is Error Handling Intro File")
	n, err := fmt.Println("Hello!")
	if err != nil {
		fmt.Println(err)

	}
	fmt.Println(n)

}
