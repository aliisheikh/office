package main

import "fmt"

func main() {

	var num int
	var sum int
	fmt.Print("Enter Any Number: ")
	fmt.Scan(&num)
	fmt.Println("Your Entered number is:", num)

	for i := 1; i <= num; i++ {
		fmt.Println("Series is: ", i, "*", i, "=", i*i)

		sum += i * i
	}
	fmt.Println("The Sum of Total Series is:", sum)
}
