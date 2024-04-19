package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("This is Recover Function File!")
	c()

}

func c() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recover in c", r)
		}
	}()
	fmt.Println("Calling Function d")
	d(0)
	fmt.Println("Returned Normally from d")

}
func d(i int) {
	if i > 5 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf(strconv.Itoa(i)))
	}
	defer fmt.Println("Defer 'Descending Order' in G:\t", i)
	fmt.Println("Printing 'Ascending' Order' in g:\t", i)
	d(i + 1)

}
