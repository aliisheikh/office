package main

import "fmt"

func main() {
	fmt.Println("This is Range Anonymous")
	c := make(chan int)

	//	send

	go func() {
		for i := 1; i <= 5; i++ {
			c <- i
		}
		close(c)
	}()

	for v := range c {
		fmt.Println(v)
	}

}
