// Topic : Interface

package main

import (
	"fmt"
	"math"
)

// 1. Structs in Go automatically implements an interface if they have all the same methods
type shape interface {
	area() float64
}

type square struct {
	length float64
}

// 2. For a struct to implements an interface, the function cannot be pointer receiver, e.g., (s *square) is not allowed
func (s square) area() float64 {
	return s.length * s.length
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func output (shp shape) {
	fmt.Printf("%T : %v\n", shp, shp.area())
}

func main() {
	s := square{length: 5}
	c := circle{radius: 5}
	
	output(s)
	output(c)
}
