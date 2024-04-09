package main

import "fmt"

func main() {
	fmt.Println("CallBack")

	s := Maths(23, 45, add)
	fmt.Println(s)

	d := Maths(90, 20, sub)
	fmt.Println(d)

	f := Maths(10, 10, mul)
	fmt.Println(f)

	g := Maths(50, 10, div)
	fmt.Println(g)

	h := Maths(60, 10, mod)
	fmt.Println(h)

}

func Maths(a int, b int, f func(int, int) int) int {
	return f(a, b)

}

func add(a int, b int) int {

	return a + b

}

func sub(a int, b int) int {
	return a - b

}

func mul(a int, b int) int {
	return a * b

}
func div(a int, b int) int {
	return a / b

}

func mod(a int, b int) int {
	return a % b

}
