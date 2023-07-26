package main

import (
	"fmt"
	"strconv"
)

type airport struct {
	name        string
	established int
	flight      []plane
	staffs      map[string]staff //ID : Staff
}

type plane struct {
	name       string
	make       string
	engine     string
	year       int
	capacity   int
	passengers []person
}

type person struct {
	name string
	age  int
}

type staff struct {
	name   string
	salary int
	years  int
}

var z airport

func main() {
	fmt.Println("Start Airport program")

	z1 := new(airport)
	fmt.Println(z1)

	z.established = 1000
	sydneyairport := new(airport)
	sydneyairport.newAirport("Sydney", 1900)

	fmt.Println("New Value:")
	fmt.Println("", *sydneyairport)

}

type UUID struct {
	counter int
}

// create new Airport given name
// since it is a method
// we will assume airport have been initialised first
func (a *airport) newAirport(s string, establish int) {

	// set airport to its zero value

	// assign name and establish date
	a.name = s
	a.established = establish

	// create staff and assign default
	if a.staffs == nil {
		a.staffs = make(map[string]staff)
	}

	// assign Bob as default staffs
	a.staffs["1"] = staff{name: "Bob", salary: 10, years: 0}

	// allocate plane in memory
	a.flight = make([]plane, 0)

	// create 2 passengers via struct literal
	person1 := person{"John", 20}
	person2 := person{"Lucy", 30}

	var arraypeople []person
	arraypeople = append(arraypeople, person1, person2)

	// create plane
	plane := plane{
		name:       "Qantas",
		make:       "Boeing",
		engine:     "Boeing",
		year:       2022,
		capacity:   200,
		passengers: arraypeople,
	}

	//
	a.flight = append(a.flight, plane)
	fmt.Println("Completed Operation", a.flight)

}

func (u *UUID) generateUUID() string {
	// UUID in the form of [8] [][][][][][][][]
	// 10^ 10 fits 10 billion
	// 1 digit of int can hold 8 bits (1 byte)
	var uuid string

	uuid = strconv.Itoa(u.counter)
	u.counter++

	return uuid
}
