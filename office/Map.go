package main

import "fmt"

func main() {

	sec15 := map[string]int{
		"Ali":    22,
		"Mateen": 24,
		"Behzad": 25,
	}
	fmt.Println(sec15)
	fmt.Println("The Age of Ali was:", sec15["Ali"])

	fmt.Println(len(sec15))

	//	Through make slice Function

	a := make(map[string]int)

	a["ALi"] = 23
	a["Mateen"] = 21

	fmt.Println("The Values of a are:", a)

	for k, v := range a {
		fmt.Println(k, v)
	}
	fmt.Println("-----------------------------")
	//It only Prints the Values
	fmt.Println("It can ONLY Print Values:")

	for _, v := range a {
		fmt.Println(v)
	}
	fmt.Println("-----------------------------")

	fmt.Println("It Only Prints the key")

	for k := range a {
		fmt.Println(k)
	}

	fmt.Println("Map with the Slice ")

	xi := []int{42, 45, 56}

	for i, v := range xi {
		fmt.Println(i, v)
	}

	//It only Prints the Values
	fmt.Println("It can ONLY Print Values:")

	for _, v := range a {
		fmt.Println(v)
	}

	fmt.Println("It Only Prints the key")

	for i := range a {
		fmt.Println(i)
	}

	delete(a, "Ali")
	fmt.Println(a["Ali"])

	//	OK comma idiom
	//for non existing value
	v, ok := a["georgey"]
	if ok {
		fmt.Println("Value is Found", v)
	} else {
		fmt.Println("Value didn't Exist")
	}

	//for existing value

	v, okk := a["ALi"]
	if okk {
		fmt.Println("Value is Found", v)
	} else {
		fmt.Println("Value didn't Exist")
	}

}
