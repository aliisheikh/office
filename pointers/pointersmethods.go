package main

import "fmt"

type dog struct {
	first string
	age   int
	kind  string
}

func (d dog) walk() {

	fmt.Println("My dog name is ", d.first, " And his Age is", d.age, "It's breed is", d.kind)
}

func (d *dog) run() {
	d.first = "Rover"

	fmt.Println("My dog name is ", d.first, " And his Age is", d.age, "It's breed is", d.kind)
}

func main() {
	fmt.Println("This is Pointer Method File")

	d1 := dog{
		first: "Henry",
		age:   5,
		kind:  "Labrador",
	}

	d2 := &dog{
		first: "patty",
		age:   4,
		kind:  "rottviller",
	}
	d1.walk()
	d1.run()
	d2.walk()
	d2.run()

}
