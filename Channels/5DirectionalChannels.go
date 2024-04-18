package main

import "fmt"

func main() {
	fmt.Println("This is Directional Channels")

	//  1
	c := make(chan int)
	//   2
	cs := make(chan<- int)
	//   3
	cr := make(<-chan int)

	fmt.Println("----------")
	fmt.Printf("%T Type of c \n", c)
	fmt.Printf("%T The Type of cr \n", cr)
	fmt.Printf("%T The Type of cs \n", cs)

	//	1
	go func() {
		c <- 10
	}()
	fmt.Println("chan int", <-c)

	//	General to specific converts
	fmt.Printf("\t<---------->\t")
	fmt.Printf("c %T \n", (<-chan int(c)))
	fmt.Printf("c %T \n", (chan<- int(c)))

}
