package main

import (
	"fmt"
	"math"
)

type calulation interface {
	area() float64
	circumference() float64
}

type cirl struct {
	radius float64
}

func (c cirl) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c cirl) circumference() float64 {
	return 2 * math.Pi * c.radius
}

// if a variable has an interface, we can call methods in the interface
// we can define a new method as measure
// to print all
func measure(c calulation) {
	fmt.Println(c)
	fmt.Println(c.area())
	fmt.Println(c.circumference())
}

func main() {
	fmt.Println("Interfaces")
	// interfaces provides a method to a Type

	var newCircle cirl
	newCircle = cirl{radius: 20}
	fmt.Println(newCircle)

	fmt.Println(newCircle.area())

	fmt.Println(newCircle.circumference())

	// since we have a measure method
	fmt.Println("Generics in interface")
	measure(newCircle)

}
