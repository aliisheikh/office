package main

import "fmt"

func main() {
	ub := make(chan string, 2)
	//ab := make(chan string, 1)
	//This will run
	ub <- "Muhammad"
	//This will give an error because there is only one space
	//Which contain 'ub <-"Muhammad"'
	ub <- "Ali"
	fmt.Println(<-ub)
	fmt.Println(<-ub)

	//fmt.Println(<-ab)

}
