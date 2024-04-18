package main

import "fmt"

func main() {
	a := 20
	fmt.Println("The Value of a is:", a)
	fmt.Println("The Memory Location a is: ", &a)
	fmt.Printf("The Type of 'a' is %T", a)
	fmt.Println("\n")
	b := "Muhammad"
	fmt.Println("The s is:", b)
	fmt.Println("The Address of 's' is:", &b)
	fmt.Printf("The Type os 's' is %T", b)

	fmt.Println("\n")

	//It will print value of b -> Muhammad
	c := &b
	fmt.Println(*c)

}
