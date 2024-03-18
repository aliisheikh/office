package main

import "fmt"

func main() {
	var percentage int

	fmt.Println("Enter marks to check the Any Subject Result: ")
	fmt.Scanf("%d", &percentage)
	fmt.Println("You entered:", percentage)

	if percentage >= 90 && percentage <= 100 {
		fmt.Println("The student got 'A' Grade")

	} else if percentage >= 101 {
		fmt.Println("Invalid")
	} else if percentage >= 80 {

		fmt.Println("The student got 'B' Grade")
	} else if percentage >= 70 {
		fmt.Println("The student got 'C' Grade")
	} else if percentage >= 60 {
		fmt.Println("The student got 'D' Grade")
	} else if percentage >= 40 && percentage >= 0 {
		fmt.Println("The student got 'E' Grade")

	} else if percentage <= 0 {
		fmt.Println("Invalid")
	}

}
