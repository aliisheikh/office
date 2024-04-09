package main

import "fmt"

type person struct {
	firrst_name string
	Last_name   string
	age         int
}

type study struct {
	person
	college    string
	university string
	graduate   bool
}

func main() {

	//Simple Struct
	p1 := person{
		firrst_name: "Muhammad",
		Last_name:   "Ali",
		age:         22,
	}

	p2 := person{
		firrst_name: "Muhammad ",
		Last_name:   "Umer",
		age:         25,
	}

	fmt.Println("The p1 is:", p1)
	fmt.Println("The p2 is:", p2)

	//	Embeded Struct

	fmt.Println("Embeded Struct:")

	st1 := study{
		person: person{
			firrst_name: "Muhammad",
			Last_name:   "Ali",
			age:         22,
		},
		college:    "PGC",
		university: "UCP",
		graduate:   false,
	}

	st2 := study{
		person: person{
			firrst_name: "Muhammad",
			Last_name:   "Umer",
			age:         25,
		},

		college:    "PGC",
		university: "PUCIT",
		graduate:   true,
	}

	fmt.Println("The st1 is:", st1)
	fmt.Println("The st2 is:", st2)

	//	Anonymous Struct

	//It is useful for single function
	fmt.Println("Anonymous Struct")

	an1 := struct {
		firrst_name string
		Last_name   string
		age         int
	}{
		firrst_name: "Muhammad",
		Last_name:   "Ali",
		age:         22,
	}

	an2 := struct {
		firrstName string
		LastName   string
		age        int
	}{
		firrstName: "Muhammad ",
		LastName:   "Umer",
		age:        25,
	}

	fmt.Println("The p1 is:", an1)
	fmt.Println("The p2 is:", an2)

	//	Composition Struct

}
