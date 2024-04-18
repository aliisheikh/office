package main

import "fmt"

func main() {

	ev := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)

	go sender(ev, odd, quit)

	receiver(ev, odd, quit)

}

//Send Channel

func sender(even, odd chan<- int, quit chan<- bool) {
	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}

	}

	close(quit)

}

func receiver(even, odd <-chan int, quit <-chan bool) {
	for {
		select {
		case v := <-even:
			fmt.Println("The Value of Even is:", v)
		case v := <-odd:
			fmt.Println("The Value of Odd is:", v)
		case i, ok := <-quit:
			if !ok {
				fmt.Println("Comma Ok Idiom:", i, ok)
			} else {
				fmt.Println("Else Part for Comma Ok idioms", i)

			}

		}
	}

}
