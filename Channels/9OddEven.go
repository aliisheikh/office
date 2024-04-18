package main

import (
	"fmt"
)

func main() {
	odd := make(chan int)
	even := make(chan int)
	quit := make(chan int)

	go send(even, odd, quit)

	receive(even, odd, quit)

}

func send(e, o, q chan<- int) {
	for i := 0; i <= 20; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}

	close(o)
	close(e)
	close(q)

}

func receive(e, o, q <-chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println("Channel from Even", v)
		case v := <-o:
			fmt.Println("Channel from Odd", v)
		case v := <-q:
			fmt.Println("Quit", v)
			return
		}
	}

}
