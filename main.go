package main

import (
	"awesomeProject/office/Animal"
	"awesomeProject/office/MyPackage"
	"fmt"
)

type canine struct {
	name string
	age  int
}

func main() {
	an := canine{
		name: "Foda",
		age:  Animal.Years(1),
	}

	fmt.Println(an)
	MyPackage.Function1()
	MyPackage.Function2()

}
