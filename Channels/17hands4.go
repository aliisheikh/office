package main

import "fmt"

func main() {

	even := make(chan int)
	odd := make(chan int)
	exit := make(chan int)

	go senders(even, odd, exit)

	receivers(even, odd, exit)

}

func senders(e, o, q chan<- int) {
	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(e)
	close(o)
	close(q)

}
func receivers(e, o, q <-chan int) {
	for {
		select {
		case v := <-e:
			fmt.Printf("Channel from even Number is\n", v)
		case v := <-o:
			fmt.Println("Channel from Odd Number is\n", v)
		case v := <-q:
			fmt.Println("Quit", v)
			return
		}

	}
}
