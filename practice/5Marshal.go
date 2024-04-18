package main

import (
	"encoding/json"
	"fmt"
)

type dog struct {
	Name  string `json:"Name"`
	Age   int    `json:"Age"`
	Breed string `json: Breed"`
}

func (d dog) String() string {
	return fmt.Sprintln("Dog Name is ", d.Name, " its Age is ", d.Age, "And it's Breed is", d.Breed)

}
func main() {
	d1 := dog{
		Name:  "Spikey",
		Age:   5,
		Breed: "Labrador",
	}

	d2 := dog{
		Name:  "Terry",
		Age:   5,
		Breed: "rototiller",
	}

	animal := []dog{d1, d2}
	fmt.Println(animal)

	bs, err := json.Marshal(animal)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Println(string(bs))
}
