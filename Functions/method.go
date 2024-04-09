package main

import "fmt"

type man struct {
	firstname string
	lastname  string
	age       int
}

func (m man) listen() {
	fmt.Println("I am", m.firstname, m.lastname)
	fmt.Println("My age is", m.age)

}

func main() {

	m1 := man{
		firstname: "Muhammad",
		lastname:  "Ali",
		age:       23,
	}

	//fmt.Println(m1)
	m2 := man{
		firstname: "Muhammad",
		lastname:  "Umer",
		age:       26,
	}
	//fmt.Println(m2)

	m1.listen()
	m2.listen()
}
