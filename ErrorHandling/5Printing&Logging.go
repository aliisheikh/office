package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Open("ExampleFile.txt")
	if err != nil {
		fmt.Println("Error Occurs", err)
		//log.Println("Error Occurs", err)
		//log.Fatal(err)
		//panic(err)

	}

}
