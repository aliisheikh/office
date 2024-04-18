package main

import "fmt"

type customErr struct {
	infoErr string
}

func (ce customErr) Error() string {
	return fmt.Sprintf("Here's the Error %v", ce.infoErr)
}

func main() {
	fmt.Println("Hello! This is Hands On Exercise 3 ")

	c1 := customErr{
		infoErr: "This is Info Error struct",
	}
	fou(c1)
}
func fou(e error) {
	fmt.Println("foo ran -", e, "\n")
}
