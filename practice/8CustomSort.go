package main

import (
	"fmt"
	"sort"
)

type identity struct {
	First string
	Last  string
	Age   int
}

func (i identity) String() string {
	return fmt.Sprintf("%s", i.First, i.Last, i.Age)

}

type ByName []identity

func (n ByName) Len() int {
	return len(n)

}
func (n ByName) Swap(a int, k int) {
	n[a], n[k] = n[k], n[a]

}
func (n ByName) Less(a, k int) bool {
	return n[a].Age < n[k].Age
}

func main() {
	fmt.Println("Custom Sort")

	i := []identity{
		{"James", "Bond", 22},
		{"Alex", "King", 12},
		{"Muhammad", "Ali", 10},
		{"Muhammad", "ALi", 21},
		{"Bilal", "Tanveer", 20},
	}
	fmt.Println(i)
	sort.Sort(ByName(i))
	fmt.Println("After Sorting By Age", i)

}
