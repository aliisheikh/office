package main

import "fmt"

type male struct {
	firstname string
	lastname  string
	age       int
}

type female struct {
	male
	firstname string
	lastname  string
	age       int
}

func (ml male) speak() {
	fmt.Println("My name is ", ml.firstname, ml.lastname)
	fmt.Println("My Age is", ml.age)
}

func (fm female) speak() {
	fmt.Println("My name is ", fm.firstname, fm.lastname)
	fmt.Println("My Age is", fm.age)

}

type human interface {
	speak()
}

func something(h human) {
	h.speak()

}
func main() {

	ml1 := male{
		firstname: "Muhammad",
		lastname:  "Ali",
		age:       23,
	}
	fm1 := female{

		firstname: "Emily",
		lastname:  "rose",
		age:       27,
	}
	//ml1.speak()
	//fm1.speak()
	something(ml1)
	something(fm1)

}

// In interface and Polymorphism:

// Values can be more than one type
