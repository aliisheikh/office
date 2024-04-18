package main

import (
	"encoding/json"
	"fmt"
)

type me struct {
	First string
	Age   int
}

func (m me) String() string {
	return fmt.Sprintln("First name and age is ", m.First, m.Age)

}
func main() {

	fmt.Println("Hands1")

	m1 := me{
		First: "Muhammad ALi",
		Age:   23,
	}
	m2 := me{
		First: "Muhammad Umer",
		Age:   26,
	}

	people := []me{m1, m2}
	fmt.Println(people)

	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))

}
