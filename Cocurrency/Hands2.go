package main

import "fmt"

type person struct {
	first string
}

func (p *person) speak() {
	fmt.Println("Hello This is Speak Function")

}

type huamn interface {
	speak()
}

func saySomething(h huamn) {
	h.speak()

}
func main() {
	p := person{
		first: "James",
	}
	//	Does Not work "saySomething (p)" because of Pointer
	saySomething(&p)
	p.speak()

}

//Receiver             Values
// -------------------------------
//(t T)                T and  *T
//(t *T)                *T
