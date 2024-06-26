package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}
type shape interface {
	area() float64
}

func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func info(s shape) {
	fmt.Println("Area", s.area())

}

func main() {
	c := circle{5}
	fmt.Println(c.area())

}

//This code does run - notice the difference -- method set
//determines the INTERFACE that the type implement
