package main

import "fmt"

func main() {

	x := sum("todd mcleod", 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println("The sum is:", x)
}

func sum(s string, ii ...int) (string int) {

	fmt.Println(ii)
	fmt.Printf("%T\n", ii)

	as := 0
	for _, v := range ii {
		as += v
	}
	fmt.Println(s)
	return as
}
