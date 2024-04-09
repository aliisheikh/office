package main

import "fmt"

// Value Semantic
func addOne(v int) int {
	return v + 2
}

func addOnep(v *int) int {
	return *v + 1
}

func main() {

	fmt.Println("Semantic Value")
	//	Semantic Value
	p := 90
	fmt.Println(p)

	fmt.Println(addOne(p))
	fmt.Println(p)

	// Semantic Pointer
	fmt.Println("Semantic Pointer")

	fmt.Println(addOnep(&p))
	fmt.Println(p)

}

// go run -gcflags -m file.go
