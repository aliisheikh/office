package main

import (
	"fmt"
	"sort"
)

func main() {

	asort := []int{2, 1, 3, 5, 4, 6, 78, 98, 76, 12, 43, 27, 1}
	sort.Ints(asort)
	fmt.Println(asort)

	//	String sort

	ssort := []string{"Apple", "Cat", "Banana", "King", "Girl"}
	sort.Strings(ssort)
	fmt.Println(ssort)
}
