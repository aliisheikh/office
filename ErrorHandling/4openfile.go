package main

import (
	"fmt"

	"os"
)

func main() {
	b, err := os.Open("File1.txt")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer b.Close()

	bs, err := os.ReadFile(b.Name())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(bs))

}
