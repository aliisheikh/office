package main

import (
	"fmt"
	"sort"
)

//Sorted By Age
//
//type man struct {
//	name string
//	age  int
//}
//
//func (m man) String() string {
//	return fmt.Sprintf("%s: %d", m.name, m.age)
//
//}
//
//type ByAge []man
//
//func (a ByAge) Len() int {
//	return len(a)
//
//}
//
//func (a ByAge) Swap(i, j int) {
//	a[i], a[j] = a[j], a[i]
//
//}
//
//func (a ByAge) Less(i, j int) bool {
//	return a[i].age < a[j].age
//
//}
//
//func main() {
//	people := []man{
//		{"James", 22},
//		{"jenny", 18},
//		{"John", 45},
//		{"Williams", 38},
//	}
//	fmt.Println(people)
//	sort.Sort(ByAge(people))
//	fmt.Println(people)
//
//}

//Sorted By Name

type male struct {
	Name string
	Age  int
}

func (l male) String() string {
	return fmt.Sprintf("%s", l.Name, l.Age)

}

type ByName []male

func (n ByName) Len() int {
	return len(n)
}

func (n ByName) Swap(a, k int) {

	n[a], n[k] = n[k], n[a]

}
func (n ByName) Less(a, k int) bool {
	return n[a].Name < n[k].Name

}

func main() {
	people := []male{
		{"James", 22},
		{"Henry", 18},
		{"Thomas", 45},
		{"Williams", 38},
	}
	fmt.Println(people)
	sort.Sort(ByName(people))
	fmt.Println(people)

}
