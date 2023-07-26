package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
	address
}

type address struct {
	location string
	number   int
	suburb   string
}

func main() {
	fmt.Println("Hello World")

	// initialise the structure
	// a is nil
	var a person

	// a  is zero value
	// allocate phase
	a = person{} // alocate empty
	b := address{location: "Jo Street", number: 20, suburb: "Sydney"}

	a = person{
		"John",
		20,
		b,
	}

	fmt.Println(a)

	var z test
	z = test{

		{input1: 2,
			input2: 2,
			result: 2},
		{
			input1: 3,
			input2: 5,
			result: 8,
		},
		{
			input1: 10,
			input2: 20,
			result: 30,
		},
	}

	for i := range z {
		fmt.Println(i)
	}
}

type test []struct {
	input1 int
	input2 int
	result int
}
