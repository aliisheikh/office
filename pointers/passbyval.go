package main

import "fmt"

func delta(n *int) {
	*n = 56

}

func strdelta(ii []string) {
	ii[0] = "Henry"

}

func main() {
	q := 21
	fmt.Println(q)
	delta(&q)
	fmt.Println(q)

	//	Pointers by slice

	sq := []string{"James", "Jenny", "john", "Amanda"}
	fmt.Println(sq)
	strdelta(sq)
	fmt.Println(sq)
}
