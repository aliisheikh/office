package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type personn struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   string `json:"Age"`
}

func main() {
	s := `s [{"First":"Marshal","Last":"JSON","Age":22},{"First":"Jenny","Last":"Williams","Age":20}]`
	bs := []byte(s)

	fmt.Printf("%T \n", s)
	fmt.Printf("%T \n", bs)

	people := []personn{}

	//var people []personn
	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Println("All the people Data", s)

	for i, v := range people {
		fmt.Println("Person Number", i)
		fmt.Println(v.First, v.Last, v.Age)

	}
	fmt.Fprintln(os.Stdout, "Hello, World: This is Fprintln ")
	io.WriteString(os.Stdout, "Hello, World: This is Write String ")
}
