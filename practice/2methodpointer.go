package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) walk() {
	fmt.Println("My first name is ", p.first, "My LAst Name is ", p.last, "And My Age is ", p.age)

}
func (p *person) run() {

	//p.first = "Sheikh"
	fmt.Println("My First name is ", p.first, "My Last Name is", p.last, "And My Age is ", p.age)
}

func main() {
	h1 := person{
		first: "Muhammad",
		last:  "Ali",
		age:   23,
	}
	h2 := person{
		first: "Muhammad",
		last:  "Umer",
		age:   26,
	}
	h1.walk()
	h1.run()

	h2.walk()
	h2.run()

}
