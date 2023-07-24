package main

import (
	"fmt"
	"time"
)

type Car struct {
	Name   string
	Old    int
	Places []string
}

// every argument in a function have its value copied
func (c Car) addPlace(s string) Car {
	c.Places = append(c.Places, s)
	return c
}

// every argument in a function have its value copied
func (c Car) changeName(s string) Car {
	c.Name = s
	return c
}

// this time we use a pointer to modify the value
func (c *Car) appendPlaces(s string) {
	c.Places = append(c.Places, s)
}

func main() {
	fmt.Println("Methods")

	const day = 24 * time.Hour
	fmt.Println(day.Seconds())

	var newCar Car
	newCar = Car{Name: "Honda", Old: 1000}

	fmt.Println(newCar)
	newCar = newCar.addPlace("Sydney")
	newCar = newCar.addPlace("Melbourne")

	fmt.Println(newCar)

	newCar = newCar.changeName("Suzuki")
	fmt.Println(newCar)

	// lets try the appendPlaces
	// it takes a pointer object which means we can modify it
	// since it relates to the direct memory address
	newCar.appendPlaces("Boston")
	fmt.Println(newCar)

}
