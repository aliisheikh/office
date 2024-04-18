package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type dogg struct {
	Name  string `json:"Name"`
	Age   int    `json:Age"`
	Breed string `json:"Breed"`
}

func main() {
	um := `[{"Name":"Spikey","Age":5,"Breed":"Labrador"},{"Name":"Terry","Age":5,"Breed":"rototiller"}]`

	fmt.Println("JSON Input", um)

	bs := []byte(`[{"Name":"Spikey","Age":5,"Breed":"Labrador"},{"Name":"Terry","Age":5,"Breed":"rototiller"}]`)
	fmt.Println(bs)

	animal := []dogg{}
	err := json.Unmarshal(bs, &animal)
	if err != nil {
		fmt.Println("There's is an Error", err)
	}
	fmt.Println(um)

	for i, v := range animal {
		fmt.Println(i)
		fmt.Println(v.Name, v.Age, v.Breed)
	}

	error := json.NewEncoder(os.Stdout).Encode(animal)
	if err != nil {
		fmt.Println(error)
	}

}
