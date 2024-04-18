package main

import "fmt"

func intdelta(n *int) {

	*n = 30

}

func strdelta(d []string) {
	d[0] = "Henry"

}

func main() {

	fmt.Println("This is Pass By Value file in folder Pointer")

	e := 45
	fmt.Println(e)
	intdelta(&e)
	//fmt.Println(&e)
	fmt.Println(e)

	f := []string{"James", "John", "Williams", "Amanda"}
	fmt.Println(f)
	strdelta(f)
	fmt.Println(f)

}
