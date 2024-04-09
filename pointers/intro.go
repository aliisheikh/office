package main

import "fmt"

func main() {
	//// Declaring a variable
	//x := 24
	//// Simply print
	//fmt.Println(x)
	////	Address of that variable
	//fmt.Println(&x)
	//fmt.Printf("%v \n %T \n", &x, &x)
	//s := "Muhammad"
	//fmt.Printf("%v \n %T \n", &s, &s)
	//
	//y := &s
	//fmt.Printf("%v \t %T \n", y, &y)
	//fmt.Println(*y)

	x := 32
	fmt.Println(x)

	fmt.Println(&x)

	fmt.Printf("%v \n %T \n", x, &x)

	y := "Ali"
	fmt.Println(y)
	fmt.Printf("%v \n %T \n", y, &y)

	//To show the Value of the Address, We use '*'.
	s := &y
	fmt.Println(*s)

}
