package main

import "fmt"

func main() {

	a := foco()
	fmt.Println(a)

	as := fujo()
	fmt.Println(as())

}

func foco() int {

	return 32
}

func fujo() func() int {
	return func() int {
		return 48

	}
}
