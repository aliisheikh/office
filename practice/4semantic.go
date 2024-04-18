package main

import "fmt"

func addval(i int) int {
	return i + 1

}
func addPointer(i *int) int {
	return *i + 2
}

func main() {

	//Semantic By Value
	b := 90
	fmt.Println(b)

	fmt.Println(addval(b))

	//	Semantic By Pointer

	a := 100
	fmt.Println(a)
	fmt.Println(addPointer(&a))
}
