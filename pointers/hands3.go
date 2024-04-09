package main

import "fmt"

type person struct {
	first string
}

func main() {
	p1 := person{
		first: "Henry",
	}
	fmt.Println(p1)
	p1 = Namechange(p1, "John")

	changenamep(&p1, "kazumi")
	fmt.Println(p1)

}

func Namechange(p person, s string) person {
	p.first = s
	return p
}

func changenamep(p *person, s string) {
	p.first = s

}
