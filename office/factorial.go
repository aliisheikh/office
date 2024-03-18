package main

import "fmt"

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	var num int
	fmt.Println("Enter the Number to find the Factorial: ")
	fmt.Scanf("%d", &num)
	fmt.Println("You entered:", num)

	result := factorial(num)
	fmt.Println("Factorial of the number is:", result)
}
