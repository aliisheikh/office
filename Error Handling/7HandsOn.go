package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First   string   `json:"First"`
	Last    string   `json:"Last"`
	Sayings []string `json:"Sayings"`
}

func main() {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken", "Stir Fry", "Apple", "PineApple"},
	}

	bs, err := json.Marshal(p1)
	if err != nil {
		// log fatal Shut the program Down
		log.Fatal("This is an Error", err)
	}
	fmt.Println(string(bs))

}
