package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {

	p1 := person{
		First: "Marshal",
		Last:  "JSON",
		Age:   22,
	}
	p2 := person{
		First: "Jenny",
		Last:  "Williams",
		Age:   20,
	}
	people := []person{p1, p2}
	fmt.Println(people)

	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))

}

//JSON will work if we do Upper Case the struct values.
