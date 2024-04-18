package main

import "fmt"

func main() {

	s := make(chan int)

	//send
	go fou(s)
	//receive
	for v := range s {
		fmt.Println(v)
	}

}

// send
func fou(c chan<- int) {
	for i := 0; i < 5; i++ {
		c <- i
	}
	close(c)
}
