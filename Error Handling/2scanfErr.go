package main

import "fmt"

func main() {
	var answer1, answer2 string

	var answer3 int
	fmt.Println("First Name")
	_, err := fmt.Scan(&answer1)
	if err != nil {
		panic(err)
	}
	fmt.Println("Last Name")
	_, err = fmt.Scan(&answer2)
	if err != nil {
		panic(err)

	}

	fmt.Println("Age")
	_, err = fmt.Scan(&answer3)
	if err != nil {
		panic(err)
	}
	fmt.Println(answer1, answer2, answer3)

}
