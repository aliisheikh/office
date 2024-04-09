package main

import "fmt"

func main() {
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	x := unfurl(xi...)
	fmt.Println("The Sum is", x)

}

func unfurl(ii ...int) int {
	fmt.Println(ii)
	fmt.Printf("%T\n", ii)

	n := 0
	for _, v := range ii {
		n += v

	}
	return n

}
