package main

import "fmt"

func main() {
	fmt.Printf("This is Hands On Exercise No.3")
	c := joo()
	received(c)

}
func received(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}

}
func joo() <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i <= 5; i++ {
			c <- i
		}
		close(c)
	}()

	return c
}
