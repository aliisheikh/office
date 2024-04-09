package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	xb, err := readFile("poem.txt")
	if err != nil {
		log.Fatalf("Error occured", err)
	}
	fmt.Println(xb)
	fmt.Println(string(xb))
}

func readFile(filename string) ([]byte, error) {

	xb, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("Error is found in this file %s", err)

	}
	return xb, nil

}
