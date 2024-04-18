package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println("Integer Sort")
	j := []int{2, 5, 7, 9, 10, 34, 23, 21, 12, 41, 76}

	fmt.Println("The Values of Array are", j)

	sort.Ints(j)
	fmt.Println("The Values After sorting is", j)

	fmt.Println("String Sort")

	k := []string{"William", "Amanda", "John", "Behzad", "Matee", "Bilal", "Umer", "Ali"}

	fmt.Println("The Values of K is", k)

	sort.Strings(k)

	fmt.Println("The Value After String Sort is", k)

	fmt.Println("Float Sort")

	l := []float64{3.44, 1.2, 3, 3.0, 1.0, 1, 10.3, 4.5, 7.1, 7.0, 6.5, 2.5}
	fmt.Println("The Float Values are", l)

	sort.Float64s(l)

	fmt.Println("The Value After float Sort is ", l)
}
