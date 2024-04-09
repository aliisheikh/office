package main

import "fmt"

type father struct {
	first string
	age   int
}

func (f father) speak() {
	fmt.Println("My father Name is", f.first, " His Age is", f.age)
}

func main() {
	xi := []int{1, 2, 3, 4, 4, 55, 5, 6, 6, 6, 77, 7}
	x := addition(xi...)
	fmt.Println(x)

	// Hands On Exercise 2
	q := barbara()
	fmt.Println(q)
	w := fuoco()
	fmt.Println(w)

	//	Hands On Exercise 3
	ds := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sd := play(ds...)
	fmt.Println(sd)

	defer fmt.Println(3)
	defer fmt.Println(2)
	defer fmt.Println(1)
	fmt.Println(0)

	p := father{
		first: "BABA",
		age:   24,
	}
	p.speak()

	//	Hands On Exercise (Anonymous Function)

	func() {
		fmt.Println("Hi! I'm An Anonymous Function")
	}()

	func(s string) {
		fmt.Println("Hi! My name is", s)
	}("Muhammad")

	nm := num()
	fmt.Println(nm())
}

func addition(ii ...int) int {
	fmt.Println(ii)

	n := 0
	for _, v := range ii {
		n += v

	}
	return n

}

//Hands On Exercise 2

func fuoco() int {
	return 23

}
func barbara() int {
	return 32

}

//Hands On Exercise 3

func play(ii ...int) int {
	fmt.Println(ii)

	n := 0
	for _, v := range ii {
		n += v

	}
	return n

	//	Hands On Exercise 4

	//Defer Function( Last In First Out)

}

// Return The function

func num() func() int {
	return func() int {
		return 42
	}

}
