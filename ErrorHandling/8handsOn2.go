package main

import (
	"errors"
	"fmt"
)

func main() {
	e := errors.New("Hello Errors!")
	fmt.Println(e)
	fmt.Printf("%T", e)

}
