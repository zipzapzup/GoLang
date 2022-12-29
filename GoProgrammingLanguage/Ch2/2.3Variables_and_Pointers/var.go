package main

import "fmt"

func main() {
	var a int = 10

	// Variable declaration
	var string1 string = "Hello"
	var int1 int = 90
	var s string

	fmt.Printf("%v, %v, %v\n", string1, int1, s)

	// Printing the memory address via & pointer
	// shows the address of x
	// var x int
	// fmt.Printf("Memory Address of x: %v", &x)
	// fmt.Printf("Value of pointer x: %v", *x)

	// Pointers
	x := 1
	p := &x
	fmt.Println(*p, p)

	*p = 0
	fmt.Println(*p)

	// re assign the value of a. You cannot re-declare.
	a = 20
	fmt.Println(a)

}
