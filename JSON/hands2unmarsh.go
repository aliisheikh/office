package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type animal struct {
	Name string `json:"Name"`
	Age  int    `json:"Age"`
}

func main() {
	um := `[{"First":"Muhammad ALi","Age":23},{"First":"Muhammad Umer","Age":26}]`
	fmt.Println(um)

	bs := []byte{}

	fmt.Printf("%T", um)
	fmt.Printf("%T", bs)

	people := []animal{}
	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("All people Data is ", um)

	for i, v := range people {
		fmt.Println("Person Number", i)
		fmt.Println(v.Name, v.Age)

	}

	//	Hands On 3

	error := json.NewEncoder(os.Stdout).Encode(people)
	if err != nil {
		fmt.Println(error)

	}
}
