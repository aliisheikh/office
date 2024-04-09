package main

import (
	"fmt"
	"sort"
)

func main() {
	//Int Sort
	zi := []int{11, 2, 3, 21, 12, 3, 4, 5, 6, 9, 8, 90}
	sort.Ints(zi)
	fmt.Println(zi)

	//String sort
	xi := []string{"Queen", "King ", "Beta", "Gamma", "Alpha", "Turban", "Airport"}
	sort.Strings(xi)
	fmt.Println(xi)
	//	Float Sort
	li := []float64{1.2, 4.3, 4.55, 2.22, 2.33, 7.6, 6.7, 8.0, 1.0, 0.99}
	sort.Float64s(li)
	fmt.Println(li)

}
